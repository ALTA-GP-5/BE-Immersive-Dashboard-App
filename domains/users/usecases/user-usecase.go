package userusecase

import (
	entity "immersive/domains/users/entities"
)

type UserUseCase struct {
	Repo entity.IrepoUser
}

func New(repo entity.IrepoUser) *UserUseCase {
	return &UserUseCase{
		Repo: repo,
	}
}

func (u *UserUseCase) Store(userEntity entity.UserEntity) {
	u.Repo.Store(userEntity)
}

func (u *UserUseCase) Update(userEntity entity.UserEntity) error {
	return u.Repo.Update(userEntity)
}

func (u *UserUseCase) Delete(userEntity entity.UserEntity) error {
	return u.Repo.Delete(userEntity)
}

func (u *UserUseCase) GetById(userEntity entity.UserEntity) ( entity.UserEntity, error ) {
	return u.Repo.FindById(userEntity)
}
