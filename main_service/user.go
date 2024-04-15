package main_service

type User struct {
	Id          int    `json:"-" db:"id"`
	TimeCreated string `json:"time_created"`
	TimeUpdated string `json:"time_updated"`

	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`

	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UserPublic struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
