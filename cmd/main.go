package main

import (
	"net/http"
	routes "practical_exercise/cmd/Routes"
	"practical_exercise/internal/configs"
	"practical_exercise/internal/databases"
	"practical_exercise/internal/databases/student"
	"practical_exercise/internal/services"

	"github.com/rs/zerolog/log"
)

func main() {
	db, err := configs.InitDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("github.com/rs/zerolog/log")
	}

	dbStore := databases.DBStore{
		Student: student.NewManagement(db),
		DBconn:  db,
	}

	serviceContainer := services.Container{
		Student: services.NewUser(dbStore),
	}

	r := routes.Routes(serviceContainer)

	log.Info().Msg("Server is starting at port 8080")
	http.ListenAndServe(":8080", r)
}
