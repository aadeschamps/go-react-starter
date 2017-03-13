package models

import (
	"database/sql"
)

// UserModel contains db and performs actions
type UserModel struct {
	Db *sql.DB
}

// User to be passed around to actions / json encoded in responses
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Create takes in a User and saves it to the database, returning the saved user
func (um *UserModel) Create(u *User) (*User, error) {
	var newUser User
	query := `
		INSERT INTO 
		users (name)
		VALUES ($1)
		RETURNING id
	`
	err := um.Db.QueryRow(query, u.Name).Scan(&newUser.ID)
	if err != nil {
		return nil, err
	}
	return um.Get(&newUser)
}

// GetAll retreives all Users from the database and returns them in a list of Users
func (um *UserModel) GetAll() ([]*User, error) {
	query := `
		SELECT *
		FROM users
	`
	rows, err := um.Db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	users := []*User{}
	for rows.Next() {
		newUser := User{}
		err := rows.Scan(&newUser.ID, &newUser.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, &newUser)
	}
	return users, nil
}

// Get gets a specific user
func (um *UserModel) Get(u *User) (*User, error) {
	query := `
		SELECT *
		FROM users
		WHERE
		id = $1
	`
	row := um.Db.QueryRow(query, u.ID)
	var user User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update takes in a user and updates it with new information
func (um *UserModel) Update(u *User) (*User, error) {
	query := `
		UPDATE users
		SET
		name = $2
		WHERE
		id = $1
	`
	_, err := um.Db.Exec(query, u.ID, u.Name)
	if err != nil {
		return nil, err
	}
	return um.Get(u)
}
