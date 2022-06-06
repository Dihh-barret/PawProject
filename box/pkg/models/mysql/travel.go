package mysql

import (
	"database/sql"
)

type TravelModel struct {
	DB *sql.DB
}
