package usermodel

import (
	entity "immersive/domains/users/entities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func EntityToModel(userEntity entity.UserEntity) User {
	return User{
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Password: userEntity.Password,
	}
}

func ModelToEntity(userModel User) entity.UserEntity {
	return entity.UserEntity{
		UID:   userModel.ID,
		Name:  userModel.Name,
		Email: userModel.Email,
	}
}
