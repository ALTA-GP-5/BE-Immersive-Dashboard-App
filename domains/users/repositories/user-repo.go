package userrepo

import (
	entity "immersive/domains/users/entities"
	model "immersive/domains/users/models"
	"immersive/exceptions"
)

type userRepo struct {
	DB []model.User
}

func New(db []model.User) *userRepo {
	return &userRepo{
		DB: db,
	}
}

func (r *userRepo) Store(userEntity entity.UserEntity) {
	id := len(r.DB) + 1
	userModel := model.EntityToModel(userEntity)
	userModel.ID = uint(id)
	r.DB = append(r.DB, userModel)
}

func (r *userRepo) Update(userEntity entity.UserEntity) error {
	userIndex := -1
	userModel := model.EntityToModel(userEntity)

	for index, product := range r.DB {
		if product.ID == userModel.ID {
			userIndex = index
		}
	}

	if userIndex == -1 {
		return exceptions.NewNotFoundError("data user not found")
	}

	return nil
}

func (r *userRepo) Delete(userEntity entity.UserEntity) error {
	userIndex := -1

	for index, product := range r.DB {
		if product.ID == userEntity.UID {
			userIndex = index
		}
	}

	if userIndex == -1 {
		return exceptions.NewNotFoundError("data user not found")
	}

	userArr1 := r.DB[0:userIndex]
	userArr2 := r.DB[userIndex+1:]

	r.DB = append(userArr1, userArr2...)

	return nil
}

func (r *userRepo) FindById(userEntity entity.UserEntity) ( entity.UserEntity, error)  {
	userIndex := -1

	for index, product := range r.DB {
		if product.ID == userEntity.UID {
			userIndex = index
		}
	}

	if userIndex == -1 {
		return entity.UserEntity{}, exceptions.NewNotFoundError("data user not found")
	}

	return model.ModelToEntity(r.DB[userIndex]), nil
}
