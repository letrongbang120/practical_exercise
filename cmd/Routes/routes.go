package routes

import (
	"encoding/json"
	"net/http"
	"practical_exercise/internal/services"

	"github.com/go-chi/chi/v5"
)

func GetApi(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("testing")
}

func Routes(service services.Container) http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/test", GetApi)
		UserRoute(r, service)
	})
	return r
}
