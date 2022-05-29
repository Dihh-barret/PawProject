package mysql

import (
	"database/sql"
	"merlin.com/box/pkg/models"
)

type TravelModel struct {
	DB *sql.DB
}
