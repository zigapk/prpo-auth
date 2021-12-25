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

// NewUserHandle  @Summary      Creates new user.
// @Description   Create new user.
// @Produce        json
// @Param         email     body      string  true  "New user's email"
// @Param         name      body      string  true  "New user's name"
// @Param         password  body      string  true  "New user's password"
// @Success       200       {object}  User
// @Failure        500  {object}  errors.ResponseError
// @Router        /users/ [post]
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

// UserHandle    @Summary      Liveness probe
// @Description  Get single user by id.
// @Produce      json
// @Param        uid  path      string  true  "User uid."
// @Success      200  {object}  User
// @Failure      500  {object}  errors.ResponseError
// @Router       /users/{uid} [get]
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

// GetUsersHandle  @Summary      Get the list of users.
// @Description    Get the list of users.
// @Produce       json
// @Success        200  {object}  []User
// @Failure       500       {object}  errors.ResponseError
// @Router         /users/ [get]
func GetUsersHandle(w http.ResponseWriter, r *http.Request) {
	var users []*user.User

	users, err := user.All()

	if err != nil {
		logger.Log.Warn().Err(err).Send()
	}

	// Write response.
	res, _ := json.Marshal(users)
	_, _ = w.Write(res)
}
