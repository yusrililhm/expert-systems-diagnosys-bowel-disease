package domain

import "time"

type User struct {
	Id        uint
	Username  string
	Email     string
	FullName  string
	Gender    bool
	Role      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
