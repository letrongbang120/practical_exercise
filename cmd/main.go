package cmd

import (
	"net/http"
	routes "practical_exercise/cmd/Routes"
	"practical_exercise/internal/configs"
	"practical_exercise/internal/databases"
	"practical_exercise/internal/databases/student"

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

	r := routes.Routes()

	log.Info().Msg("Server is starting at port 8080")
	http.ListenAndServe(":8080", r)
}
