package routes

import (
	"encoding/json"
	"net/http"
	"practical_exercise/internal/services"

	"github.com/go-chi/chi/v5"
)

func testUserApi(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("user api worked !!!")
}

func UserRoute(r chi.Router, service services.Container) {
	r.Route("/user", func(r chi.Router) {
		r.Get("/test", testUserApi)
	})
}
