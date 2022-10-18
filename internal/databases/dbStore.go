package databases

import (
	"database/sql"
	"practical_exercise/internal/databases/student"
)

type DBStore struct {
	Student student.Repository
	DBconn  *sql.DB
}
