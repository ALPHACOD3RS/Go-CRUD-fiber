package main

import "gorm.io/gorm"


type UserRoles string

const (
	AdminRole UserRoles = "admin"
	UserRole UserRoles =  "user"
	GuestRole UserRoles = "Guesr"
)
type User struct {
	gorm.Model
	Name 		string 		`json:"name" gorm:"not null"`
	Email 		string		`json:"email" gorm:"unique; not null"`
	Password 	string 		`json:"password" gorm:"not null"`
	Role 		string		`json:"role" gorm:"default:'user'"`
	Posts 		[]Post		`json:"posts" gorm:"foreignkey:UserID"`
}

type Post struct {
	gorm.Model
	Title   string     `json:"title"`
	Content string	   `json:"content"`
	UserID  string     `json:"user_id"`
}