package loginusecase

import (
	entity "immersive/domains/login/entities"
	"immersive/exceptions"
	"immersive/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	Repo entity.ILoginRepo
}

func New(repo entity.ILoginRepo) *loginUsecase {
	return &loginUsecase{
		Repo: repo,
	}
}

func (u *loginUsecase) Login(mentor entity.MentorEntity) (string, error) {
	mentorResult, err := u.Repo.GetByEmail(mentor)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(mentorResult.Password), []byte(mentor.Password))

	if err != nil {
		return "", exceptions.NewBadRequestError("email or password not match!")
	}

	token, err := middlewares.CreateToken(mentorResult.MentorID, mentorResult.Role)
	if err != nil {
		return "", exceptions.NewInternalServerError(err.Error())
	}

	return token, nil
}
