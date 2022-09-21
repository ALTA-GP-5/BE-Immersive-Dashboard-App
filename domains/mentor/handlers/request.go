package mentorhandler

import entity "immersive/domains/mentor/entities"

type request struct {
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Team     string `json:"team" form:"team" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Status   string `json:"status" form:"status" validate:"required"`
}

func requestToEntity(mentor request) entity.MentorEntity {
	return entity.MentorEntity{
		FullName: mentor.FullName,
		Email:    mentor.Email,
		Team:     mentor.Team,
		Role:     mentor.Role,
		Password: mentor.Password,
		Status:   mentor.Status,
	}
}
