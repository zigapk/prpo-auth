package handle

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zigapk/prpo-auth/internal/handle/errors"
	"github.com/zigapk/prpo-auth/internal/logger"
	middleware "github.com/zigapk/prpo-auth/internal/middleware"
	"github.com/zigapk/prpo-auth/internal/models/user"
	"github.com/zigapk/prpo-auth/internal/util"
)

type newUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type userEditRequest struct {
	Name string `json:"name"`
}

// NewUserHandle creates new user.
func NewUserHandle(w http.ResponseWriter, r *http.Request) {
	// Get data from body.
	newReq := &newUserRequest{}
	if !util.ParseJSON(w, r, newReq) {
		return
	}

	// Create new user.
	u, err := user.New(newReq.Email, newReq.Name, newReq.Password)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Write response.
	res, _ := json.Marshal(u)
	_, _ = w.Write(res)
}

// UserHandle handles getting a single user.
func UserHandle(w http.ResponseWriter, r *http.Request) {
	// Get authorized user and check their permissions.
	authUser := middleware.UserFromRequest(r)

	// Get requested user.
	uid := mux.Vars(r)["uid"]

	if authUser.UID != uid {
		errors.Response(w, errors.InsufficientPermissions)
		return
	}

	u, err := user.ForUID(uid)
	if err != nil {
		logger.Log.Warn().Err(err).Send()

		errors.Response(w, errors.DatabaseError)
		return
	}

	// Write response.
	res, _ := json.Marshal(u)
	_, _ = w.Write(res)
}
