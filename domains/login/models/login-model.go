package loginmodel

import (
	entity "immersive/domains/login/entities"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Fullname string
	Email    string
	Team     string
	Role     string
	Password string
}

func ModelToEntity(mentorModel Mentor) entity.MentorEntity {
	return entity.MentorEntity{
		MentorID: mentorModel.ID,
		Role:     mentorModel.Role,
		Email:    mentorModel.Email,
		Password: mentorModel.Password,
	}
}
