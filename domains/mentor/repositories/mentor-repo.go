package mentorrepo

import (
	entity "immersive/domains/mentor/entities"
	model "immersive/domains/mentor/models"
	"immersive/exceptions"

	"gorm.io/gorm"
)

type mentorRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *mentorRepo {
	return &mentorRepo{
		DB: db,
	}
}

func (r *mentorRepo) Insert(mentor entity.MentorEntity) error {
	dataModel := model.EntityToModel(mentor)

	tx := r.DB.Model(&model.Mentor{}).Create(&dataModel)
	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}
	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("failed to create mentor")
	}

	return nil
}

func (r *mentorRepo) GetAll(mentor entity.MentorEntity) ([]entity.MentorEntity, error) {
	dataModel := []model.Mentor{}

	tx := r.DB.Model(&model.Mentor{})

	if mentor.GeneralSearch != "" {
		tx.Where("fullname LIKE ?", "%"+mentor.GeneralSearch+"%").Or("email LIKE ?", "%"+mentor.GeneralSearch+"%").Or("team LIKE ?", "%"+mentor.GeneralSearch+"%").Or("role LIKE ?", "%"+mentor.GeneralSearch+"%").Or("mentors.status LIKE ?", "%"+mentor.GeneralSearch+"%")
	}

	if mentor.ClassStatus != "" {
		tx.Joins("JOIN classes ON classes.mentor_id = mentors.id AND classes.status = ?", mentor.ClassStatus)
	}

	tx.Find(&dataModel)

	if tx.Error != nil {
		return []entity.MentorEntity{}, exceptions.NewInternalServerError(tx.Error.Error())
	}

	var mentorEntity []entity.MentorEntity
	for _, data := range dataModel {
		mentorEntity = append(mentorEntity, model.ModelToEntity(data))
	}

	return mentorEntity, nil
}

func (r *mentorRepo) GetById(mentor entity.MentorEntity) (entity.MentorEntity, error) {
	dataModel := model.Mentor{}
	dataModel.ID = mentor.MentorID

	tx := r.DB.Model(&model.Mentor{}).First(&dataModel)
	if tx.Error != nil {
		return entity.MentorEntity{}, exceptions.NewInternalServerError(tx.Error.Error())
	}

	return model.ModelToEntity(dataModel), nil
}

func (r *mentorRepo) Update(mentor entity.MentorEntity) error {
	dataModel := model.EntityToModel(mentor)
	dataModel.ID = mentor.MentorID

	tx := r.DB.Model(&model.Mentor{}).Where("id = ?", dataModel.ID).Updates(&dataModel)
	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("failed to update mentor")
	}

	return nil
}

func (r *mentorRepo) Delete(mentor entity.MentorEntity) error {
	dataModel := model.EntityToModel(mentor)
	dataModel.ID = mentor.MentorID

	tx := r.DB.Model(&model.Mentor{}).Where("id = ?", dataModel.ID).Delete(&dataModel)
	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("failed to delete mentor")
	}

	return nil
}
