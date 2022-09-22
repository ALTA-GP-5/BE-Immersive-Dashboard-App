package menteeusecase

import (
	"errors"
	entity "immersive/domains/mentee/entities"
	"immersive/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	menteeEntity := entity.MenteeEntity{
		ClassID:         1,
		Name:            "anthony stark",
		Address:         "jl.kapitayan no.22",
		HomeAddress:     "jl.kapitayan no.22",
		Email:           "anthonystark@gmail.com",
		Gender:          true,
		Telegram:        "t.me/application",
		Phone:           "01283012",
		EmergencyName:   "Howard Stark",
		EmergencyPhone:  "12151253",
		EmergencyStatus: "Orang Tua",
		Type:            false,
		Major:           "Bachelor",
		Graduate:        "Massachussets Institute School",
		Status:          "active",
		Category:        "Informatics",
	}

	t.Run("Create mentee success", func(t *testing.T) {
		repo := new(mocks.MenteeRepoMock)
		repo.On("Insert", mock.Anything).Return(nil).Once()

		menteeusecase := New(repo)
		err := menteeusecase.Create(menteeEntity)
		assert.NoError(t, err)
		assert.NotEqual(t, err, menteeEntity)

		repo.AssertExpectations(t)
	})

	t.Run("Create mentee failed", func(t *testing.T) {
		repo := new(mocks.MenteeRepoMock)
		repo.On("Insert", mock.Anything).Return(errors.New("bad request")).Once()

		menteeusecase := New(repo)
		err := menteeusecase.Create(menteeEntity)
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	menteeEntity := entity.MenteeEntity{
		ClassID:         1,
		Name:            "anthony stark",
		Address:         "jl.kapitayan no.22",
		HomeAddress:     "jl.kapitayan no.22",
		Email:           "anthonystark@gmail.com",
		Gender:          true,
		Telegram:        "t.me/application",
		Phone:           "01283012",
		EmergencyName:   "Howard Stark",
		EmergencyPhone:  "12151253",
		EmergencyStatus: "Orang Tua",
		Type:            false,
		Major:           "Bachelor",
		Graduate:        "Massachussets Institute School",
		Status:          "active",
		Category:        "Informatics",
	}

	t.Run("Update mentee success", func(t *testing.T) {
		repo := new(mocks.MenteeRepoMock)
		repo.On("Update", mock.Anything).Return(nil).Once()

		menteeusecase := New(repo)
		err := menteeusecase.Update(menteeEntity)
		assert.NoError(t, err)
		assert.NotEqual(t, err, menteeEntity)

		repo.AssertExpectations(t)
	})

	t.Run("Update mentee failed", func(t *testing.T) {
		repo := new(mocks.MenteeRepoMock)
		repo.On("Update", mock.Anything).Return(errors.New("bad request")).Once()

		menteeusecase := New(repo)
		err := menteeusecase.Update(menteeEntity)
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.MenteeRepoMock)

	t.Run("Delete mentee success", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()

		menteeusecase := New(repo)
		err := menteeusecase.Delete(entity.MenteeEntity{})
		assert.NoError(t, err)
		assert.NotEqual(t, entity.MenteeEntity{}, err)

		repo.AssertExpectations(t)
	})

	t.Run("Delete mentee failed", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New("bad request")).Once()

		menteeusecase := New(repo)
		err := menteeusecase.Delete(entity.MenteeEntity{})
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	menteeEntity := []entity.MenteeEntity{
		{
			MenteeID: 1,
			Name:     "jhon doe2",
			Class:    "BE 11",
			Status:   "active",
			Category: "IT",
			Gender:   true,
		},
	}

	t.Run("Find all mentees success", func(t *testing.T) {
		repo := new(mocks.MenteeRepoMock)
		repo.On("FindAll", mock.Anything).Return(menteeEntity, nil).Once()

		menteeusecase := New(repo)
		result, err := menteeusecase.GetAll(entity.MenteeEntity{})
		assert.NoError(t, err)
		assert.Equal(t, result, menteeEntity)
		repo.AssertExpectations(t)
	})

	t.Run("Find all internal server error", func(t *testing.T) {
		repo := new(mocks.MenteeRepoMock)
		repo.On("FindAll", mock.Anything).Return([]entity.MenteeEntity{}, errors.New("internal server error")).Once()

		menteeusecase := New(repo)
		result, err := menteeusecase.GetAll(entity.MenteeEntity{})
		assert.Error(t, err)
		assert.NotEqual(t, result, err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	menteeEntity := entity.MenteeEntity{
		MenteeID:        1,
		Name:            "anthony stark",
		Address:         "jl.kapitayan no.22",
		HomeAddress:     "jl.kapitayan no.22",
		Email:           "anthonystark@gmail.com",
		Gender:          false,
		Telegram:        "t.me/application",
		Phone:           "01283012",
		EmergencyName:   "Howard Stark",
		EmergencyPhone:  "12151253",
		EmergencyStatus: "Orang Tua",
		Type:            true,
		Major:           "Bachelor",
		Graduate:        "Massachussets Institute School",
		Status:          "active",
		Class:           "BE 11",
		Category:        "Informatics",
	}

	t.Run("Read mentee by id success", func(t *testing.T) {
		repo := new(mocks.MenteeRepoMock)
		repo.On("Find", mock.Anything).Return(menteeEntity, nil).Once()

		menteeusecase := New(repo)
		result, err := menteeusecase.Get(entity.MenteeEntity{})
		assert.NoError(t, err)
		assert.Equal(t, result, menteeEntity)
		repo.AssertExpectations(t)
	})

	t.Run("Read by id internal server error", func(t *testing.T) {
		repo := new(mocks.MenteeRepoMock)
		repo.On("Find", mock.Anything).Return(entity.MenteeEntity{}, errors.New("internal server error")).Once()

		classusecase := New(repo)
		result, err := classusecase.Get(entity.MenteeEntity{})
		assert.Error(t, err)
		assert.NotEqual(t, result, err.Error())
		repo.AssertExpectations(t)
	})
}
