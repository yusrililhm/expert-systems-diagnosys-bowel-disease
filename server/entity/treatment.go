package entity

import "gorm.io/gorm"

type Treatment struct {
	gorm.Model
	Description string     `gorm:"not null"`
	Diseases    []*Disease `gorm:"many2many:disease_treatments"`
}
