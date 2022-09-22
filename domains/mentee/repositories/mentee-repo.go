package menteerepo

import (
	entity "immersive/domains/mentee/entities"
	model "immersive/domains/mentee/models"
	"immersive/exceptions"

	"gorm.io/gorm"
)

type menteeRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *menteeRepo {
	return &menteeRepo{
		DB: db,
	}
}

func (r *menteeRepo) Insert(menteeEntity entity.MenteeEntity) error {
	menteeModel, menteeDetailModel := model.EntityToModel(menteeEntity)

	tx := r.DB.Model(&model.Mentee{}).Create(&menteeModel)
	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	menteeDetailModel.MenteeID = menteeModel.ID
	tx2 := r.DB.Model(&model.MenteeDetail{}).Create(&menteeDetailModel)
	if tx2.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("failed to create mentee rows affected 0")
	}

	return nil
}

func (r *menteeRepo) Update(menteeEntity entity.MenteeEntity) error {
	menteeModel, menteeDetailModel := model.EntityToModel(menteeEntity)
	tx := r.DB.Model(&model.Mentee{}).Select("*").Where("id = ?", menteeModel.ID).Updates(&menteeModel)

	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	tx2 := r.DB.Model(&model.MenteeDetail{}).Select("*").Where("mentee_id = ?", menteeModel.ID).Updates(&menteeDetailModel)
	if tx2.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("failed to update mentee rows affected 0")
	}

	return nil
}

func (r *menteeRepo) Delete(menteeEntity entity.MenteeEntity) error {
	menteeModel := model.Mentee{}
	menteeModel.ID = menteeEntity.MenteeID
	tx := r.DB.Model(&model.Mentee{}).Delete(&menteeModel)

	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("failed to Delete mentee rows affected 0")
	}

	return nil
}

func (r *menteeRepo) FindAll(menteeEntity entity.MenteeEntity) ([]entity.MenteeEntity, error) {
	var menteeModelList []model.Mentee

	tx := r.DB.Model(&model.Mentee{}).Preload("MenteeDetail").Preload("Class")

	if menteeEntity.GeneralSearch != "" {
		tx.Where("name LIKE ?", "%"+menteeEntity.GeneralSearch+"%")
	}

	if menteeEntity.Status != "" {
		tx.Where("status = ?", menteeEntity.Status)
	}

	if menteeEntity.ClassID != 0 {
		tx.Where("class_id = ?", menteeEntity.ClassID)
	}

	if menteeEntity.Category != "" {
		tx.Where("category = ?", menteeEntity.Category)
	}

	tx.Find(&menteeModelList)

	if tx.Error != nil {
		return []entity.MenteeEntity{}, exceptions.NewInternalServerError(tx.Error.Error())
	}

	var menteeEntityList []entity.MenteeEntity
	for _, menteeModel := range menteeModelList {
		menteeEntityList = append(menteeEntityList, model.ModelToEntity(menteeModel))
	}

	return menteeEntityList, nil
}

func (r *menteeRepo) Find(menteeEntity entity.MenteeEntity) (entity.MenteeEntity, error) {
	menteeModel := model.Mentee{}
	menteeModel.ID = menteeEntity.MenteeID

	tx := r.DB.Model(&model.Mentee{}).Preload("MenteeDetail").Preload("Class").First(&menteeModel)

	if tx.Error != nil {
		return entity.MenteeEntity{}, exceptions.NewInternalServerError(tx.Error.Error())
	}

	return model.ModelToEntity(menteeModel), nil
}
