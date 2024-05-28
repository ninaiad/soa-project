package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"soa-main/internal/user"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (time_created, time_updated, username, password_hash, name, surname, birthday, email, phone) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", usersTable)
	row := a.db.QueryRow(query, user.TimeCreated, user.TimeUpdated, user.Username, user.Password, user.Name, user.Surname, user.Birthday, user.Email, user.Phone)
	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (a *AuthPostgres) GetUser(username, password string) (user.User, error) {
	var user user.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := a.db.Get(&user, query, username, password)
	return user, err
}

func (a *AuthPostgres) GetUserLogin(userId int) (user.User, error) {
	var user user.User
	query := fmt.Sprintf("SELECT username FROM %s WHERE id=$1", usersTable)
	err := a.db.Get(&user, query, userId)
	return user, err
}

func (a *AuthPostgres) GetUserData(userId int) (user.UserPublic, error) {
	user := user.UserPublic{}
	query := fmt.Sprintf("SELECT name, surname, birthday, email, phone FROM %s WHERE id = $1", usersTable)
	err := a.db.Get(&user, query, userId)
	return user, err
}

func (a *AuthPostgres) UpdateUser(userId int, update user.UserPublic, timeUpdated string) error {
	query := fmt.Sprintf("UPDATE %s ul SET name=$1, surname=$2, birthday=$3, email=$4, phone=$5, time_updated=$6 WHERE ul.id = $7", usersTable)
	_, err := a.db.Exec(query, update.Name, update.Surname, update.Birthday, update.Email, update.Phone, timeUpdated, userId)
	return err
}
