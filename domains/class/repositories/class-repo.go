package classrepo

import (
	entity "immersive/domains/class/entities"
	model "immersive/domains/class/models"
	"immersive/exceptions"

	"gorm.io/gorm"
)

type classRepo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *classRepo {
	return &classRepo{
		DB: db,
	}
}

func (r *classRepo) Insert(classEntity entity.ClassEntity) error {
	classModel := model.EntityToModel(classEntity)

	tx := r.DB.Model(&model.Class{}).Create(&classModel)

	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("class create failed")
	}

	return nil
}

func (r *classRepo) Update(classEntity entity.ClassEntity) error {
	classModel := model.EntityToModel(classEntity)
	classModel.ID = classEntity.ClassID

	tx := r.DB.Model(&model.Class{}).Where("id = ?", classModel.ID).Updates(&classModel)

	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("class update failed")
	}

	return nil
}

func (r *classRepo) Delete(classEntity entity.ClassEntity) error {
	classModel := model.EntityToModel(classEntity)
	classModel.ID = classEntity.ClassID

	tx := r.DB.Model(&model.Class{}).Delete(&classModel)

	if tx.Error != nil {
		return exceptions.NewInternalServerError(tx.Error.Error())
	}

	if tx.RowsAffected < 1 {
		return exceptions.NewInternalServerError("class delete failed")
	}

	return nil
}

func (r *classRepo) FindAll(classEntity entity.ClassEntity) ([]entity.ClassEntity, error) {
	var classModelList []model.Class

	tx := r.DB.Model(&model.Class{}).Preload("Mentor")

	if tx.Error != nil {
		return []entity.ClassEntity{}, exceptions.NewInternalServerError(tx.Error.Error())
	}

	if classEntity.GeneralSearch != "" {
		tx.Where("name = ?", classEntity.GeneralSearch).Or("status = ?", classEntity.GeneralSearch)
	}

	tx.Find(&classModelList)

	var entityClass []entity.ClassEntity
	for _, classModel := range classModelList {
		entityClass = append(entityClass, model.ModelToEntity(classModel))
	}

	return entityClass, nil
}

func (r *classRepo) Find(classEntity entity.ClassEntity) (entity.ClassEntity, error) {
	var classModel model.Class
	classModel.ID = classEntity.ClassID

	tx := r.DB.Model(&model.Class{}).Preload("Mentor").First(&classModel)

	if tx.Error != nil {
		return entity.ClassEntity{}, exceptions.NewInternalServerError(tx.Error.Error())
	}

	return model.ModelToEntity(classModel), nil
}
