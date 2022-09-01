package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rejvban/macro-api/internal/auth"
	"github.com/rejvban/macro-api/internal/health"
)

func New() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", health.HandleHealth).Methods(http.MethodGet)

	auth.InitRouter(router)

	return router
}
