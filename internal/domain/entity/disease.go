package entity

import "gorm.io/gorm"

type Disease struct {
	gorm.Model
	Name        string       `gorm:"not null"`
	Description string       `gorm:"not null"`
	Symptomps   []*Symptom   `gorm:"many2many:disease_symptoms"`
	Treatments  []*Treatment `gorm:"many2many:disease_treatments"`
}
