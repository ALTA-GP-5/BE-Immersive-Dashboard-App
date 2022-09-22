package feedbackrepo

import (
	entity "immersive/domains/feedback/entities"
	model "immersive/domains/feedback/models"
	"immersive/exceptions"

	"gorm.io/gorm"
)

type feedbackRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *feedbackRepo {
	return &feedbackRepo{
		DB: db,
	}
}

func (r *feedbackRepo) Insert(feedbackEntity entity.FeedBackEntity) error {
	feedbackModel := model.EntityToModel(feedbackEntity)

	tx := r.DB.Model(&model.FeedBack{}).Create(&feedbackModel)

	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("class create failed")
	}

	return nil
}

func (r *feedbackRepo) FindAll(feedbackEntity entity.FeedBackEntity) ([]entity.FeedBackEntity, error) {
	var feedbackModelList []model.FeedBack

	tx := r.DB.Model(&model.FeedBack{}).Where("mentee_id = ?", feedbackEntity.MenteeID).Preload("Mentor").Find(&feedbackModelList)

	if tx.Error != nil {
		return []entity.FeedBackEntity{}, exceptions.NewInternalServerError(tx.Error.Error())
	}

	var feedbackEntityList []entity.FeedBackEntity
	for _, feedbackModel := range feedbackModelList {
		feedbackEntityList = append(feedbackEntityList, model.ModelToEntity(feedbackModel))
	}

	return feedbackEntityList, nil
}

func (r *feedbackRepo) UpdateMentee(feedBackEntity entity.FeedBackEntity) error {
	tx := r.DB.Model(&model.Mentee{}).Where("id = ?", feedBackEntity.MenteeID).Update("status", feedBackEntity.Status)
	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("failed to update mentee")
	}

	return nil
}
