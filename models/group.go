package models

import (
	"database/sql"
)

// GroupModel contains db and performs actions
type GroupModel struct {
	Db *sql.DB
}

// Group respresents a group from the db
type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
