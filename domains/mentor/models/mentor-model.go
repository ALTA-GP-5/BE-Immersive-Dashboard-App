package mentormodel

import "gorm.io/gorm"

type Mentor struct {
	gorm.Model
	Fullname string
	Email    string
	Team     string
	Role     string
	Password string
}
