package models

import (
	"time"
)

// user struct
type User struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name" gorm:"type:varchar(100);not null"`
	Email     string     `json:"email" gorm:"unique;type:varchar(200);not null"`
	Password  string     `json:"password" gorm:"type:varchar(200);not null"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// user login struct
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
