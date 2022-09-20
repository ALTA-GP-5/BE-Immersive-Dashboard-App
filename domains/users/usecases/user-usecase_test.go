package userusecase

import (
	"errors"
	"immersive/mocks"
	"testing"

	entity "immersive/domains/users/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	t.Run("success update", func(t *testing.T) {
		repo := new(mocks.UserRepoMock)

		repo.On("Update", mock.Anything).Return(nil).Once()

		userUsecase := New(repo)

		err := userUsecase.Update(entity.UserEntity{
			UID:      1,
			Name:     "budi",
			Email:    "budi@gmail.comd",
			Password: "asdasdasd",
		})

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update", func(t *testing.T) {
		repo := new(mocks.UserRepoMock)

		repo.On("Update", mock.Anything).Return(errors.New("failed update")).Once()

		userUsecase := New(repo)

		err := userUsecase.Update(entity.UserEntity{
			UID:      1,
			Name:     "budi",
			Email:    "budi@gmail.comd",
			Password: "asdasdasd",
		})

		assert.Error(t, err)
		repo.AssertExpectations(t)
	})
}
