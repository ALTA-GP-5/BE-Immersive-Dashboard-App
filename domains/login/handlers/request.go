package loginhandler

import (
	entity "immersive/domains/login/entities"
)

type request struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func requestToEntity(mentor request) entity.MentorEntity {
	return entity.MentorEntity{
		Email:    mentor.Email,
		Password: mentor.Password,
	}
}
