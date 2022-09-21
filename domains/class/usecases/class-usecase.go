package classusecase

import (
	entity "immersive/domains/class/entities"
)

type classUsecase struct {
	Repo entity.IClassRepo
}

func New(repo entity.IClassRepo) *classUsecase {
	return &classUsecase{
		Repo: repo,
	}
}

func (u *classUsecase) Create(classEntity entity.ClassEntity) error {
	return u.Repo.Insert(classEntity)
}

func (u *classUsecase) Update(classEntity entity.ClassEntity) error {
	return u.Repo.Update(classEntity)
}

func (u *classUsecase) Delete(classEntity entity.ClassEntity) error {
	return u.Repo.Delete(classEntity)
}

func (u *classUsecase) GetAll(classEntity entity.ClassEntity) ([]entity.ClassEntity, error) {
	return u.Repo.FindAll(classEntity)
}

func (u *classUsecase) Get(classEntity entity.ClassEntity) (entity.ClassEntity, error) {
	return u.Repo.Find(classEntity)
}
