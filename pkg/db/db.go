package db

import (
	"template/Template/config"

	"github.com/Asliddin3/Template/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
)
func ConnectToDb(cfg config.Config)