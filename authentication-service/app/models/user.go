package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      RoleEnum  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserCreated struct {
	ID       int64    `json:"id"`
	Username string   `json:"username"`
	Role     RoleEnum `json:"role"`
}

type RoleEnum string

const (
	RoleAdmin RoleEnum = "admin"
	RoleUser  RoleEnum = "user"
)
