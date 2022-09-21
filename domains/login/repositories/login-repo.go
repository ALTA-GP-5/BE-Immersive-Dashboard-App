package loginrepo

import (
	entity "immersive/domains/login/entities"
	model "immersive/domains/login/models"
	"immersive/exceptions"

	"gorm.io/gorm"
)

type loginRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *loginRepo {
	return &loginRepo{
		DB: db,
	}
}

func (r *loginRepo) GetByEmail(mentor entity.MentorEntity) (entity.MentorEntity, error) {
	mentorModel := model.Mentor{}
	tx := r.DB.Model(new(model.Mentor)).Where("email = ?", mentor.Email).First(&mentorModel)

	if tx.Error != nil {
		return entity.MentorEntity{}, exceptions.NewInternalServerError(tx.Error.Error())
	}

	if mentorModel.ID == 0 {
		return entity.MentorEntity{}, exceptions.NewNotFoundError("email or password not match!")
	}

	return model.ModelToEntity(mentorModel), nil
}
