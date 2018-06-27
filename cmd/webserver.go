package cmd

import (
	"github.com/ducnt114/testprj/cmd/handlers"
	"github.com/ducnt114/testprj/drivers/mongo"
	"github.com/gorilla/mux"
	"net/http"
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

var MongoConnection *mongo.MongoConnection

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

var basePath = "/v1/file"
var routes []Route

func initRoutes() {
	routes = Routes{
		Route{
			Name:    "Ping",
			Method:  http.MethodGet,
			Pattern: basePath + "/ping",
			Handler: &handlers.PingHandler{},
		},
		Route{
			Name:    "Upload file to s3",
			Method:  http.MethodPost,
			Pattern: basePath + "/s3-upload",
			Handler: &handlers.S3UploadHandler{
				MongoConn: MongoConnection,
			},
		},
	}
}
