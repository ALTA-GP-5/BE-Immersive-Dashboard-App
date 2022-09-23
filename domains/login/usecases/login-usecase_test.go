package loginusecase

import (
	"errors"
	"immersive/mocks"
	"testing"

	entity "immersive/domains/login/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	mentorEntity := entity.MentorEntity{
		MentorID: 1,
		Role:     "admin",
		Email:    "yoasobi@gmail.com",
		Password: "$2a$12$hrKZNCUx2h5lb0WlFSTqOeQyH7EVqXkrEcDrvx1uN6WJBEruo/wgq",
	}

	t.Run("login success", func(t *testing.T) {
		repo := new(mocks.LoginRepoMock)
		repo.On("GetByEmail", mock.Anything).Return(mentorEntity, nil)

		loginusecase := New(repo)
		mentorEntity.Password = "yoasobikakeru"
		token, _, err := loginusecase.Login(mentorEntity)

		assert.NoError(t, err)
		assert.NotEqual(t, "", token)

		repo.AssertExpectations(t)
	})

	t.Run("login failed password not match", func(t *testing.T) {
		repo := new(mocks.LoginRepoMock)
		mentorEntity.Password = "$2a$12$hrKZNCUx2h5lb0WlFSTqOeQyH7EVqXkrEcDrvx1uN6WJBEruo/wgq"
		repo.On("GetByEmail", mock.Anything).Return(mentorEntity, nil)

		loginusecase := New(repo)
		mentorEntity.Password = "nande"
		token, _, err := loginusecase.Login(mentorEntity)

		assert.Equal(t, "email or password not match!", err.Error())
		assert.Equal(t, "", token)

		repo.AssertExpectations(t)
	})

	t.Run("Login user id or role empty", func(t *testing.T) {
		repo := new(mocks.LoginRepoMock)
		mentorEntity.MentorID = 0
		mentorEntity.Role = ""
		mentorEntity.Password = "$2a$12$hrKZNCUx2h5lb0WlFSTqOeQyH7EVqXkrEcDrvx1uN6WJBEruo/wgq"

		repo.On("GetByEmail", mock.Anything).Return(mentorEntity, nil)

		loginusecase := New(repo)
		mentorEntity.Password = "yoasobikakeru"
		token, _, err := loginusecase.Login(mentorEntity)

		assert.Equal(t, "empty response", err.Error())
		assert.Equal(t, "", token)

		repo.AssertExpectations(t)
	})

	t.Run("Login internal server error", func(t *testing.T) {
		repo := new(mocks.LoginRepoMock)
		repo.On("GetByEmail", mock.Anything).Return(entity.MentorEntity{}, errors.New("internal server error"))

		loginusecase := New(repo)
		token, _, err := loginusecase.Login(mentorEntity)

		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, "", token)

		repo.AssertExpectations(t)
	})
}
