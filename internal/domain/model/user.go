package model

import "time"

type User struct {
	Id        int
	Username  string
	FullName  string
	Phone     string
	BirthDate time.Time
	Gender    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
