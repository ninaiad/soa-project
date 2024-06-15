package db

import (
	"fmt"
	"log"

	"gateway/internal/user"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func runDBMigration(cfg Config) error {
	m, err := migrate.New("file://./internal/db/migration",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = runDBMigration(cfg); err != nil {
		return nil, err
	}

	log.Println(`database is up and the migrations complete.`)
	return db, nil
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user user.User) (int64, error) {
	var id int64
	query := fmt.Sprintf(`
		INSERT INTO %s
			(time_created, time_updated, username, password_hash, name, surname, birthday, email, phone)
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`,
		usersTable)
	row := a.db.QueryRow(query,
		user.TimeCreated,
		user.TimeUpdated,
		user.Username,
		user.Password,
		user.Name,
		user.Surname,
		user.Birthday,
		user.Email,
		user.Phone,
	)

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (a *AuthPostgres) GetUserId(username, password string) (int64, error) {
	var userId int64
	query := fmt.Sprintf(
		`SELECT id FROM %s WHERE username=$1 AND password_hash=$2`,
		usersTable)
	err := a.db.Get(&userId, query, username, password)
	return userId, err
}

func (a *AuthPostgres) GetUserData(userId int64) (user.UserPublic, error) {
	user := user.UserPublic{}
	query := fmt.Sprintf(
		`SELECT username, name, surname, birthday, email, phone FROM %s WHERE id = $1`,
		usersTable)
	err := a.db.Get(&user, query, userId)
	return user, err
}

func (a *AuthPostgres) UpdateUser(userId int64, update user.UserPublic, timeUpdated string) error {
	query := fmt.Sprintf(`
		UPDATE %s ul
			SET username=$1, name=$2, surname=$3, birthday=$4, email=$5, phone=$6, time_updated=$7
		WHERE ul.id = $8`,
		usersTable)
	_, err := a.db.Exec(query,
		update.Username,
		update.Name,
		update.Surname,
		update.Birthday,
		update.Email,
		update.Phone,
		timeUpdated,
		userId,
	)
	return err
}
