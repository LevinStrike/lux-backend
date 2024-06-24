package core

import "database/sql"

type Databases struct {
	Type string
	SQL  *sql.DB
}
