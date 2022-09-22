package feedbackusecase

import (
	entity "immersive/domains/feedback/entities"
	"time"
)

type feedbackUsecase struct {
	Repo entity.IFeedBackRepo
}

func New(repo entity.IFeedBackRepo) *feedbackUsecase {
	return &feedbackUsecase{
		Repo: repo,
	}
}

func (u *feedbackUsecase) Create(feedbackEntity entity.FeedBackEntity) error {

	feedbackEntity.Date = time.Now()

	err := u.Repo.Insert(feedbackEntity)
	if err != nil {
		return err
	}

	return u.Repo.UpdateMentee(feedbackEntity)
}

func (u *feedbackUsecase) GetAll(feedbackEntity entity.FeedBackEntity) ([]entity.FeedBackEntity, error) {
	return u.Repo.FindAll(feedbackEntity)
}
