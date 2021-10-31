package router

import (
	"net/http"

	"github.com/gorilla/mux"
	middleware "github.com/zigapk/prpo-auth/internal/middleware"
)

func addRoute(router *mux.Router, r Route) {
	if r.GET != nil {
		router.Handle(r.Path, middleware.AddMiddleware(r.GET, r.AuthorizedOnly, r.CustomContentType)).
			Methods("GET").
			Name(r.Name)
	}

	if r.PUT != nil {
		router.Handle(r.Path, middleware.AddMiddleware(r.PUT, r.AuthorizedOnly, r.CustomContentType)).
			Methods("PUT").
			Name(r.Name)
	}

	if r.POST != nil {
		router.Handle(r.Path, middleware.AddMiddleware(r.POST, r.AuthorizedOnly, r.CustomContentType)).
			Methods("POST").
			Name(r.Name)
	}

	if r.DELETE != nil {
		router.Handle(r.Path, middleware.AddMiddleware(r.DELETE, r.AuthorizedOnly, r.CustomContentType)).
			Methods("DELETE").
			Name(r.Name)
	}

	// Add CORS handle for OPTIONS request
	router.Handle(r.Path, middleware.AddMiddleware(http.HandlerFunc(middleware.CorsHandler), false, false)).
		Methods("OPTIONS")
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/asdfasdf", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("otle neki dela")
	// })

	for _, r := range apiRoutes() {
		addRoute(router, r)
	}

	return router
}
