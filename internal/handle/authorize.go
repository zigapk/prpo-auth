package handle

import (
	"encoding/json"
	"net/http"

	"github.com/zigapk/prpo-auth/internal/config"
	"github.com/zigapk/prpo-auth/internal/handle/errors"
	"github.com/zigapk/prpo-auth/internal/logger"
	"github.com/zigapk/prpo-auth/internal/models/user"
	"github.com/zigapk/prpo-auth/internal/util"
)

type authorizationType string

var (
	authorizationTypePassword authorizationType = "password"
	authorizationTypeToken    authorizationType = "token"
)

type authorizationRequest struct {
	Type authorizationType `json:"type"`

	Email    string `json:"email"`
	Password string `json:"password"`

	RefreshToken string `json:"refresh_token"`
}

type User struct {
	UID string
}

func FromToken(accessToken string) (*User, error) {
	claims, err := util.Validate(accessToken)
	if err != nil {
		return nil, err
	}

	u := &User{
		UID: claims.Subject,
	}
	return u, nil
}

func AuthorizeHandle(w http.ResponseWriter, r *http.Request) {

	authRequest := &authorizationRequest{}
	if !util.ParseJSON(w, r, authRequest) {
		return
	}

	switch authRequest.Type {
	case authorizationTypePassword:
		authorizeWithPassword(w, authRequest)
	case authorizationTypeToken:
		authorizeWithToken(w, authRequest)
	}
}

func authorizeWithPassword(w http.ResponseWriter, authRequest *authorizationRequest) {
	// Get user with email and password.
	u, err := user.AuthorizeWithPassword(authRequest.Email, authRequest.Password)

	// Handle errors.
	if err == user.ErrInvalidCredentials {
		errors.Response(w, errors.InvalidCredentials)
		return
	}

	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Create new refresh token.
	dev, err := u.NewAuthorizedDevice()
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Create new access token.
	token, err := u.CreateAccessToken()
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.LoginError)
		return
	}

	// Write response.
	res, _ := json.Marshal(map[string]interface{}{
		"refresh_token":    dev.Token,
		"access_token":     token,
		"access_token_ttl": config.Login.TokenTtl(),
	})
	_, _ = w.Write(res)
}

// authorizeWithToken authorizes user with provided refresh token.
// It generates new access token, and writes it to response writer.
func authorizeWithToken(w http.ResponseWriter, authRequest *authorizationRequest) {
	// Get user with email and password.
	u, err := user.AuthorizeWithToken(authRequest.RefreshToken)

	// Handle errors.
	if err == user.ErrInvalidCredentials {
		errors.Response(w, errors.InvalidCredentials)
		return
	}

	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Create new access token.
	token, err := u.CreateAccessToken()
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.LoginError)
		return
	}

	// Write response.
	res, _ := json.Marshal(map[string]interface{}{
		"access_token":     token,
		"access_token_ttl": config.Login.TokenTtl(),
	})
	_, _ = w.Write(res)
}
