package mentormodel

import (
	entity "immersive/domains/mentor/entities"

	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	Fullname string
	Email    string
	Team     string
	Role     string
	Password string
	Status   string
}

func EntityToModel(mentorEntity entity.MentorEntity) Mentor {
	return Mentor{
		Fullname: mentorEntity.FullName,
		Email:    mentorEntity.Email,
		Team:     mentorEntity.Team,
		Role:     mentorEntity.Role,
		Password: mentorEntity.Password,
		Status:   mentorEntity.Status,
	}
}

func ModelToEntity(mentor Mentor) entity.MentorEntity {
	return entity.MentorEntity{
		MentorID: mentor.ID,
		FullName: mentor.Fullname,
		Email:    mentor.Email,
		Team:     mentor.Team,
		Role:     mentor.Role,
		Status:   mentor.Status,
	}
}
