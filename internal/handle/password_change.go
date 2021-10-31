package handle

import (
	"net/http"

	"github.com/zigapk/prpo-auth/internal/handle/errors"
	"github.com/zigapk/prpo-auth/internal/logger"
	middleware "github.com/zigapk/prpo-auth/internal/middleware"
	"github.com/zigapk/prpo-auth/internal/models/user"
	"github.com/zigapk/prpo-auth/internal/util"
)

type passwordChangeRequest struct {
	UserID      string `json:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// PasswordChangeHandle handles changing user password.
func PasswordChangeHandle(w http.ResponseWriter, r *http.Request) {
	// Get data from request.
	passRequest := &passwordChangeRequest{}
	if !util.ParseJSON(w, r, passRequest) {
		return
	}

	// Get authenticated user.
	authUser := middleware.UserFromRequest(r)

	// User can only check their own password.
	// Check if authenticated user is the user for which the request was made.
	if authUser.UID != passRequest.UserID {
		errors.Response(w, errors.InvalidCredentials)
		return
	}

	// Get user object from database.
	u, err := user.ForUID(authUser.UID)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Check if old password is correct.
	if !u.PasswordMatch(passRequest.OldPassword) {
		errors.Response(w, errors.InvalidOldPassword)
		return
	}

	// Set new password.
	err = user.SetPassword(u.UID, passRequest.NewPassword)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
