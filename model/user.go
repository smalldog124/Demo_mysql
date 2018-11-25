package model

import "time"

type User struct {
	UserID      string
	FristName   string
	LastName    string
	Address     string
	PhoneNumber string
	CreatedTime time.Time
	UpdatedTime time.Time
}
