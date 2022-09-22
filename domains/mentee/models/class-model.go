package classmodel

import (
	entity "immersive/domains/mentee/entities"
	"time"

	"gorm.io/gorm"
)

type Mentee struct {
	gorm.Model
	ClassID  uint
	Name     string
	Gender   bool
	Status   string
	Category string

	Class        Class        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	MenteeDetail MenteeDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type MenteeDetail struct {
	gorm.Model
	MenteeID        uint
	Address         string
	HomeAddress     string
	Email           string
	Telegram        string
	Phone           string
	EmergencyPhone  string
	EmergencyName   string
	EmergencyStatus string
	Type            bool
	Major           string
	Graduate        string
}

type Class struct {
	gorm.Model
	MentorID  uint
	Name      string
	StartDate time.Time
	EndDate   time.Time
	Status    string
}

func EntityToModel(menteeEntity entity.MenteeEntity) (Mentee, MenteeDetail) {
	menteeModel := Mentee{
		ClassID:  menteeEntity.ClassID,
		Name:     menteeEntity.Name,
		Gender:   menteeEntity.Gender,
		Status:   menteeEntity.Status,
		Category: menteeEntity.Category,
	}
	menteeModel.ID = menteeEntity.MenteeID

	menteeDetailModel := MenteeDetail{
		MenteeID:        menteeEntity.MenteeID,
		Address:         menteeEntity.Address,
		HomeAddress:     menteeEntity.HomeAddress,
		Email:           menteeEntity.Email,
		Telegram:        menteeEntity.Telegram,
		Phone:           menteeEntity.Phone,
		EmergencyPhone:  menteeEntity.EmergencyPhone,
		EmergencyName:   menteeEntity.EmergencyName,
		EmergencyStatus: menteeEntity.EmergencyStatus,
		Type:            menteeEntity.Type,
		Major:           menteeEntity.Major,
		Graduate:        menteeEntity.Graduate,
	}

	return menteeModel, menteeDetailModel
}

func ModelToEntity(menteeModel Mentee) entity.MenteeEntity {
	return entity.MenteeEntity{
		MenteeID:        menteeModel.ID,
		ClassID:         menteeModel.ClassID,
		Name:            menteeModel.Name,
		Address:         menteeModel.MenteeDetail.Address,
		HomeAddress:     menteeModel.MenteeDetail.HomeAddress,
		Email:           menteeModel.MenteeDetail.Email,
		Gender:          menteeModel.Gender,
		Telegram:        menteeModel.MenteeDetail.Telegram,
		Phone:           menteeModel.MenteeDetail.Phone,
		EmergencyName:   menteeModel.MenteeDetail.EmergencyName,
		EmergencyPhone:  menteeModel.MenteeDetail.EmergencyPhone,
		EmergencyStatus: menteeModel.MenteeDetail.EmergencyStatus,
		Type:            menteeModel.MenteeDetail.Type,
		Major:           menteeModel.MenteeDetail.Major,
		Graduate:        menteeModel.MenteeDetail.Graduate,
		Status:          menteeModel.Status,
		Class:           menteeModel.Class.Name,
		Category:        menteeModel.Category,
	}
}
