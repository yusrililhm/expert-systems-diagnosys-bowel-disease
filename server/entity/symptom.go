package entity

import "gorm.io/gorm"

type Symptom struct {
	gorm.Model
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Question    string     `gorm:"not null"`
	Diseases    []*Disease `gorm:"many2many:disease_symptoms"`
}
