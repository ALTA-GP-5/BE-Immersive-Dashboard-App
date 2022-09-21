package classmodel

import (
	"time"

	entity "immersive/domains/class/entities"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	MentorID  uint
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Status    string

	Mentor Mentor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type Mentor struct {
	gorm.Model
	Fullname string
	Email    string
	Team     string
	Role     string
	Password string
	Status   string
}

func EntityToModel(classEntity entity.ClassEntity) Class {
	return Class{
		MentorID:  classEntity.MentorID,
		Name:      classEntity.Name,
		StartDate: classEntity.StartDate,
		EndDate:   classEntity.EndDate,
		Status:    classEntity.Status,
	}
}

func ModelToEntity(classModel Class) entity.ClassEntity {
	return entity.ClassEntity{
		ClassID:    classModel.ID,
		MentorID:   classModel.MentorID,
		Name:       classModel.Name,
		Status:     classModel.Status,
		StartDate:  classModel.StartDate,
		EndDate:    classModel.EndDate,
		MentorName: classModel.Mentor.Fullname,
	}
}
