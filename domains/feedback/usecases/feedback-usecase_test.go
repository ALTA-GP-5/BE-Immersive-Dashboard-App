package feedbackusecase

import (
	"errors"
	entity "immersive/domains/feedback/entities"
	"immersive/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	t.Run("Failed check image extension error", func(t *testing.T) {
		repo := new(mocks.FeedBackRepoMock)

		feedbackusecase := New(repo)
		err := feedbackusecase.Create(entity.FeedBackEntity{
			FileName: "",
		})
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed check image size error", func(t *testing.T) {
		repo := new(mocks.FeedBackRepoMock)

		feedbackusecase := New(repo)
		err := feedbackusecase.Create(entity.FeedBackEntity{
			FileName: "alterraDoc.pdf",
			FileSize: int64(10097153),
		})
		assert.Error(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	feedbackEntity := []entity.FeedBackEntity{
		{
			FeedBackID: 1,
			Status:     "Continue To Section 2",
			MentorName: "Mail",
			Desc:       "lorem ipsum dolor sit amet conssectur elite",
			CreatedAt:  time.Now(),
			Url:        "http://example.com",
		},
	}

	t.Run("Find all feedbacks success", func(t *testing.T) {
		repo := new(mocks.FeedBackRepoMock)
		repo.On("FindAll", mock.Anything).Return(feedbackEntity, nil).Once()

		feedbackusecase := New(repo)
		result, err := feedbackusecase.GetAll(entity.FeedBackEntity{})
		assert.NoError(t, err)
		assert.Equal(t, result, feedbackEntity)
		repo.AssertExpectations(t)
	})

	t.Run("Find all internal server error", func(t *testing.T) {
		repo := new(mocks.FeedBackRepoMock)
		repo.On("FindAll", mock.Anything).Return([]entity.FeedBackEntity{}, errors.New("internal server error")).Once()

		feedbackusecase := New(repo)
		result, err := feedbackusecase.GetAll(entity.FeedBackEntity{})
		assert.Error(t, err)
		assert.NotEqual(t, result, err.Error())
		repo.AssertExpectations(t)
	})
}
