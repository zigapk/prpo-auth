package router

import (
	"net/http"
)

type Route struct {
	Name                   string
	Path                   string
	AuthorizedOnly         bool
	CustomContentType      bool
	GET, PUT, POST, DELETE http.Handler
}
