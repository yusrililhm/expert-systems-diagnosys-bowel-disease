package model

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birth_date"`
	Gender    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
