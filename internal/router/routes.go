package router

import (
	"net/http"

	"github.com/zigapk/prpo-auth/internal/handle"
)

func apiRoutes() []Route {
	return []Route{
		{
			Name: "authorize",
			Path: "/authorize",
			POST: http.HandlerFunc(handle.AuthorizeHandle),
		},
		{
			Name:           "change_password",
			Path:           "/change_password",
			AuthorizedOnly: true,
			PUT:            http.HandlerFunc(handle.PasswordChangeHandle),
		},
		{
			Name:           "users",
			Path:           "/users",
			AuthorizedOnly: false,
			POST:           http.HandlerFunc(handle.NewUserHandle),
			GET:            http.HandlerFunc(handle.GetUsersHandle),
		},
		{
			Name:           "user",
			Path:           "/users/{uid}",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.UserHandle),
		},
		{
			Name:              "signing_key",
			Path:              "/signing_key",
			CustomContentType: true,
			GET:               http.HandlerFunc(handle.SigningKeyHandle),
		},
	}
}
