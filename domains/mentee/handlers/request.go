package handlerclass

import (
	entity "immersive/domains/mentee/entities"
)

type request struct {
	ClassID         uint   `json:"class_id" form:"class_id" validate:"required"`
	Name            string `json:"name" form:"name" validate:"required"`
	Address         string `json:"address" form:"address" validate:"required"`
	HomeAddress     string `json:"home_address" form:"home_addres" validate:"required"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	Gender          bool   `json:"gender" form:"gender"`
	Telegram        string `json:"telegram" form:"telegram" validate:"required"`
	Phone           string `json:"phone" form:"phone" validate:"required"`
	EmergencyName   string `json:"emergency_name" form:"emergency_name" validate:"required"`
	EmergencyPhone  string `json:"emergency_phone" form:"emergency_phone" validate:"required"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status" validate:"required"`
	Type            bool   `json:"type" form:"type"`
	Major           string `json:"major" form:"major" validate:"required"`
	Graduate        string `json:"graduate" form:"graduate" validate:"required"`
	Status          string `json:"status" form:"status" validate:"required"`
	Category        string `json:"category" form:"category" validate:"required"`
}

func requestToEntity(menteeRequest request) entity.MenteeEntity {
	return entity.MenteeEntity{
		ClassID:         menteeRequest.ClassID,
		Name:            menteeRequest.Name,
		Email:           menteeRequest.Email,
		Address:         menteeRequest.Address,
		HomeAddress:     menteeRequest.HomeAddress,
		Telegram:        menteeRequest.Telegram,
		Phone:           menteeRequest.Phone,
		EmergencyName:   menteeRequest.EmergencyName,
		EmergencyPhone:  menteeRequest.EmergencyPhone,
		EmergencyStatus: menteeRequest.EmergencyStatus,
		Gender:          menteeRequest.Gender,
		Type:            menteeRequest.Type,
		Major:           menteeRequest.Major,
		Graduate:        menteeRequest.Graduate,
		Status:          menteeRequest.Status,
		Category:        menteeRequest.Category,
	}
}
