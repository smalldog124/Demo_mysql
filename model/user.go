package model

import "time"

type User struct {
	UserID      string    `json:"user_id" db:"user_id"`
	FristName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Address     string    `json:"addess" db:"addess"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	CreatedTime time.Time `json:"created_time" db:"created_time"`
	UpdatedTime time.Time `json:"updated_time" db:"updated_time"`
}

type UserInfo struct {
	User []User `json:"user"`
}

type ResponseNewUser struct {
	UserID int64 `json:"user_id"`
}
