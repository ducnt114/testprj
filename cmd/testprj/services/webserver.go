package services

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ducnt114/testprj/cmd/testprj/services/handlers"
)

// Route --
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}

// Routes --
type Routes []Route

// NewRouter --
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	initRoutes()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}

var basePath = "/v1/chat"
var routes []Route

func initRoutes() {
	routes = Routes{
		Route{
			Name:    "Upload file to s3",
			Method:  http.MethodPost,
			Pattern: basePath + "/s3-upload",
			Handler: &handlers.S3UploadHandler{},
		},
	}
}
