package classmodel

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	MentorID  uint
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Status    string

	Mentor Mentor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Mentor struct {
	gorm.Model
	Fullname string
	Email    string
	Team     string
	Role     string
	Password string
	Status   string
}
