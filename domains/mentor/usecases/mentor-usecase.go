package mentorusecase

import (
	entity "immersive/domains/mentor/entities"
	"immersive/exceptions"

	"golang.org/x/crypto/bcrypt"
)

type mentorUsecase struct {
	Repo entity.IMentorRepo
}

func New(repo entity.IMentorRepo) *mentorUsecase {
	return &mentorUsecase{
		Repo: repo,
	}
}

func (u *mentorUsecase) Create(mentor entity.MentorEntity) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mentor.Password), bcrypt.DefaultCost)
	if err != nil {
		return exceptions.NewInternalServerError(err.Error())
	}

	mentor.Password = string(hashedPassword)

	return u.Repo.Insert(mentor)
}

func (u *mentorUsecase) ReadAll(mentor entity.MentorEntity) ([]entity.MentorEntity, error) {
	return u.Repo.GetAll(mentor)
}

func (u *mentorUsecase) ReadById(mentor entity.MentorEntity) (entity.MentorEntity, error) {
	return u.Repo.GetById(mentor)
}

func (u *mentorUsecase) Update(mentor entity.MentorEntity) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mentor.Password), bcrypt.DefaultCost)
	if err != nil {
		return exceptions.NewInternalServerError(err.Error())
	}

	mentor.Password = string(hashedPassword)

	return u.Repo.Update(mentor)
}

func (u *mentorUsecase) Delete(mentor entity.MentorEntity) error {
	return u.Repo.Delete(mentor)
}
