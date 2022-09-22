package mentorusecase

import (
	"errors"
	"testing"

	entity "immersive/domains/mentor/entities"
	"immersive/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mentorEntity := entity.MentorEntity{
		FullName: "jhon doe",
		Email:    "jhondoe@gmail.com",
		Team:     "academic",
		Role:     "admin",
		Password: "jhondoe",
		Status:   "active",
	}

	t.Run("Create mentor success", func(t *testing.T) {
		repo := new(mocks.MentorRepoMock)
		repo.On("Insert", mock.Anything).Return(nil).Once()

		mentorusecase := New(repo)
		err := mentorusecase.Create(mentorEntity)
		assert.NoError(t, err)
		assert.NotEqual(t, err, mentorEntity)

		repo.AssertExpectations(t)
	})

	t.Run("Create mentor failed", func(t *testing.T) {
		repo := new(mocks.MentorRepoMock)
		repo.On("Insert", mock.Anything).Return(errors.New("bad request")).Once()

		mentorusecase := New(repo)
		err := mentorusecase.Create(mentorEntity)
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}

func TestReadAll(t *testing.T) {
	mentorEntity := []entity.MentorEntity{
		{
			MentorID: 1,
			FullName: "jhon doe",
			Email:    "jhondoe@gmail.com",
			Team:     "academic",
			Role:     "admin",
			Status:   "active",
		},
	}

	t.Run("Read all mentors success", func(t *testing.T) {
		repo := new(mocks.MentorRepoMock)
		repo.On("GetAll", mock.Anything).Return(mentorEntity, nil).Once()

		mentorusecase := New(repo)
		result, err := mentorusecase.ReadAll(entity.MentorEntity{})
		assert.NoError(t, err)
		assert.Equal(t, result, mentorEntity)
		repo.AssertExpectations(t)
	})

	t.Run("Read all internal server error", func(t *testing.T) {
		repo := new(mocks.MentorRepoMock)
		repo.On("GetAll", mock.Anything).Return([]entity.MentorEntity{}, errors.New("internal server error")).Once()

		mentorusecase := New(repo)
		result, err := mentorusecase.ReadAll(entity.MentorEntity{})
		assert.Error(t, err)
		assert.NotEqual(t, result, err.Error())
		repo.AssertExpectations(t)
	})
}

func TestReadById(t *testing.T) {
	mentorEntity := entity.MentorEntity{
		FullName: "jhon doe",
		Email:    "jhondoe@gmail.com",
		Team:     "academic",
		Role:     "admin",
		Status:   "active",
	}

	t.Run("Read mentor by id success", func(t *testing.T) {
		repo := new(mocks.MentorRepoMock)
		repo.On("GetById", mock.Anything).Return(mentorEntity, nil).Once()

		mentorusecase := New(repo)
		result, err := mentorusecase.ReadById(entity.MentorEntity{})
		assert.NoError(t, err)
		assert.Equal(t, result, mentorEntity)
		repo.AssertExpectations(t)
	})

	t.Run("Read by id internal server error", func(t *testing.T) {
		repo := new(mocks.MentorRepoMock)
		repo.On("GetById", mock.Anything).Return(entity.MentorEntity{}, errors.New("internal server error")).Once()

		mentorusecase := New(repo)
		result, err := mentorusecase.ReadById(entity.MentorEntity{})
		assert.Error(t, err)
		assert.NotEqual(t, result, err.Error())
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mentorEntity := entity.MentorEntity{
		FullName: "jhon doe",
		Email:    "jhondoe@gmail.com",
		Team:     "academic",
		Role:     "admin",
		Status:   "graduated",
	}

	t.Run("Update mentor success", func(t *testing.T) {
		repo := new(mocks.MentorRepoMock)
		repo.On("Update", mock.Anything).Return(nil).Once()

		mentorusecase := New(repo)
		err := mentorusecase.Update(mentorEntity)
		assert.NoError(t, err)
		assert.NotEqual(t, err, mentorEntity)

		repo.AssertExpectations(t)
	})

	t.Run("Update mentor failed", func(t *testing.T) {
		repo := new(mocks.MentorRepoMock)
		repo.On("Update", mock.Anything).Return(errors.New("bad request")).Once()

		mentorusecase := New(repo)
		err := mentorusecase.Update(mentorEntity)
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.MentorRepoMock)

	t.Run("Delete mentor success", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()

		mentorusecase := New(repo)
		err := mentorusecase.Delete(entity.MentorEntity{})
		assert.NoError(t, err)
		assert.NotEqual(t, entity.MentorEntity{}, err)

		repo.AssertExpectations(t)
	})

	t.Run("Delete mentor failed", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New("bad request")).Once()

		mentorusecase := New(repo)
		err := mentorusecase.Delete(entity.MentorEntity{})
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}
