package menteeusecase

import (
	entity "immersive/domains/mentee/entities"
)

type menteeUsecase struct {
	Repo entity.IMenteeRepo
}

func New(repo entity.IMenteeRepo) *menteeUsecase {
	return &menteeUsecase{
		Repo: repo,
	}
}

func (u *menteeUsecase) Create(menteeEntity entity.MenteeEntity) error {
	return u.Repo.Insert(menteeEntity)
}

func (u *menteeUsecase) Update(menteeEntity entity.MenteeEntity) error {
	return u.Repo.Update(menteeEntity)
}

func (u *menteeUsecase) Delete(menteeEntity entity.MenteeEntity) error {
	return u.Repo.Delete(menteeEntity)
}

func (u *menteeUsecase) GetAll(menteeEntity entity.MenteeEntity) ([]entity.MenteeEntity, error) {
	return u.Repo.FindAll(menteeEntity)
}

func (u *menteeUsecase) Get(menteeEntity entity.MenteeEntity) (entity.MenteeEntity, error) {
	return u.Repo.Find(menteeEntity)
}
