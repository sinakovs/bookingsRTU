package dbrepo

import (
	"database/sql"

	"github.com/sinakovs/bookings/internal/config"
	"github.com/sinakovs/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresREpo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
