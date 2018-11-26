package model

import "time"

type User struct {
	UserID      string    `json:"user_id"`
	FristName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Address     string    `json:"addess"`
	PhoneNumber string    `json:"phone_number"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

type UserInfo struct {
	User []User `json:"user"`
}
