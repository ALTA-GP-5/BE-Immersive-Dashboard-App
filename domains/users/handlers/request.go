package userhandler

import (
	entity "immersive/domains/users/entities"
)

type Request struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func RequestToEntity(request Request) entity.UserEntity {
	return entity.UserEntity{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
