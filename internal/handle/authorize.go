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

type authorizeAccessToken struct {
	AccessToken    string `json:"access_token"`
	AccessTokenTtl int    `json:"access_token_ttl"`
}

type authorizeRefreshToken struct {
	RefreshToken    string `json:"refresh_token"`
	AccessToken     string `json:"access_token"`
	AccessTokenTtl  int    `json:"access_token_ttl"`
	RefreshTokenTtl int    `json:"refresh_token_ttl"`
}

type User struct {
	UID string
}

// AuthorizeHandle  @Summary      Authorize either using a JWT token or a username/password combo.
// @Description     Authorize either using a JWT token or a username/password combo
// @Produce         application/json
// @Param           type      body      string  true   "Either `token` or `password`."
// @Param           email     body      string  false  "Users email in case type is set to `password`."
// @Param           password  body      string  false  "Users password in case type is set to `password`."
// @Param           token     body      string  false  "Refresh token in case type is set to `token`."
// @Success         200       {object}  authorizeRefreshToken
// @Failure         500       {object}  errors.ResponseError
// @Router          /authorize [post]
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
	tokenRefresh, err := u.CreateRefreshToken()
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.LoginError)
		return
	}
	dev, err := u.NewAuthorizedDevice(tokenRefresh)
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
	res, _ := json.Marshal(authorizeRefreshToken{
		RefreshToken:    dev.Token,
		RefreshTokenTtl: config.Login.RefreshTokenTtl(),
		AccessToken:     token,
		AccessTokenTtl:  config.Login.TokenTtl(),
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
	res, _ := json.Marshal(authorizeAccessToken{
		AccessToken:    token,
		AccessTokenTtl: config.Login.TokenTtl(),
	})
	_, _ = w.Write(res)
}
