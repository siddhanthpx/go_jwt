package models

type User struct {
	ID       int
	Name     string
	Email    string `gorm:"unique"`
	Password string
}
