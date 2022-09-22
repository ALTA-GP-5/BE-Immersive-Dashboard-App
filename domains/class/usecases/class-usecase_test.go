package classusecase

import (
	"errors"
	entity "immersive/domains/class/entities"
	"immersive/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	classEntity := entity.ClassEntity{
		MentorID: 1,
		Name:     "Backend Engineer Batch 11",
		Status:   "ongoing",
	}

	t.Run("Create class success", func(t *testing.T) {
		repo := new(mocks.ClassRepoMock)
		repo.On("Insert", mock.Anything).Return(nil).Once()

		classusecase := New(repo)
		err := classusecase.Create(classEntity)
		assert.NoError(t, err)
		assert.NotEqual(t, err, classEntity)

		repo.AssertExpectations(t)
	})

	t.Run("Create class failed", func(t *testing.T) {
		repo := new(mocks.ClassRepoMock)
		repo.On("Insert", mock.Anything).Return(errors.New("bad request")).Once()

		classusecase := New(repo)
		err := classusecase.Create(classEntity)
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	classEntity := entity.ClassEntity{
		MentorID: 1,
		Name:     "Backend Engineer Batch 11",
		Status:   "ongoing",
	}

	t.Run("Update class success", func(t *testing.T) {
		repo := new(mocks.ClassRepoMock)
		repo.On("Update", mock.Anything).Return(nil).Once()

		classusecase := New(repo)
		err := classusecase.Update(classEntity)
		assert.NoError(t, err)
		assert.NotEqual(t, err, classEntity)

		repo.AssertExpectations(t)
	})

	t.Run("Update class failed", func(t *testing.T) {
		repo := new(mocks.ClassRepoMock)
		repo.On("Update", mock.Anything).Return(errors.New("bad request")).Once()

		classusecase := New(repo)
		err := classusecase.Update(classEntity)
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.ClassRepoMock)

	t.Run("Delete class success", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()

		classusecase := New(repo)
		err := classusecase.Delete(entity.ClassEntity{})
		assert.NoError(t, err)
		assert.NotEqual(t, entity.ClassEntity{}, err)

		repo.AssertExpectations(t)
	})

	t.Run("Delete class failed", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New("bad request")).Once()

		classusecase := New(repo)
		err := classusecase.Delete(entity.ClassEntity{})
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))

		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	classEntity := []entity.ClassEntity{
		{
			MentorID:   1,
			MentorName: "Fakhry",
			Name:       "BE BATCH 11",
			Status:     "ongoing",
		},
	}

	t.Run("Find all classes success", func(t *testing.T) {
		repo := new(mocks.ClassRepoMock)
		repo.On("FindAll", mock.Anything).Return(classEntity, nil).Once()

		classusecase := New(repo)
		result, err := classusecase.GetAll(entity.ClassEntity{})
		assert.NoError(t, err)
		assert.Equal(t, result, classEntity)
		repo.AssertExpectations(t)
	})

	t.Run("Find all internal server error", func(t *testing.T) {
		repo := new(mocks.ClassRepoMock)
		repo.On("FindAll", mock.Anything).Return([]entity.ClassEntity{}, errors.New("internal server error")).Once()

		classusecase := New(repo)
		result, err := classusecase.GetAll(entity.ClassEntity{})
		assert.Error(t, err)
		assert.NotEqual(t, result, err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGet(t *testing.T) {
	classEntity := entity.ClassEntity{
		MentorID:   1,
		MentorName: "Fakhry",
		Name:       "BE BATCH 11",
		Status:     "ongoing",
	}

	t.Run("Read class by id success", func(t *testing.T) {
		repo := new(mocks.ClassRepoMock)
		repo.On("Find", mock.Anything).Return(classEntity, nil).Once()

		classusecase := New(repo)
		result, err := classusecase.Get(entity.ClassEntity{})
		assert.NoError(t, err)
		assert.Equal(t, result, classEntity)
		repo.AssertExpectations(t)
	})

	t.Run("Read by id internal server error", func(t *testing.T) {
		repo := new(mocks.ClassRepoMock)
		repo.On("Find", mock.Anything).Return(entity.ClassEntity{}, errors.New("internal server error")).Once()

		classusecase := New(repo)
		result, err := classusecase.Get(entity.ClassEntity{})
		assert.Error(t, err)
		assert.NotEqual(t, result, err.Error())
		repo.AssertExpectations(t)
	})
}
