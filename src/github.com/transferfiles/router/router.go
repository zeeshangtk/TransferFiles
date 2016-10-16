package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type routerConfig struct {
	Methods     []string
	URL         string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Transfer file is up and running"))
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
			HandlerFunc: healthCheckHandler,
		},
	}
	return &routers
}
