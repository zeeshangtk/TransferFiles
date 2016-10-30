package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/transferfiles/handlers"
)

type routerConfig struct {
	Methods     []string
	URL         string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	routes := getRoutes()

	for _, route := range *routes {
		router.HandleFunc(route.URL, route.HandlerFunc).Methods(route.Methods...)
	}
	return router
}

func getRoutes() *[]routerConfig {
	routers := []routerConfig{
		routerConfig{
			Methods:     []string{"GET"},
			URL:         "/healthcheck",
			HandlerFunc: handlers.HealthCheckHandler,
		},
		routerConfig{
			Methods:     []string{"POST"},
			URL:         "/receiveFiles",
			HandlerFunc: handlers.ReceiverHandler,
		},
		routerConfig{
			Methods:     []string{"POST"},
			URL:         "/sendFiles",
			HandlerFunc: handlers.SendFileHandler,
		},
	}
	return &routers
}
