package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
)

func GetApi(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("testing")
}

func Routes() http.Handler {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Get("/", GetApi)
	})
	return r
}
