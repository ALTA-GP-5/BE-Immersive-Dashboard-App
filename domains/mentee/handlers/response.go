package handlerclass

import (
	entity "immersive/domains/mentee/entities"
)

type response struct {
	MenteeID        uint   `json:"id"`
	ClassID         uint   `json:"class_id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	HomeAddress     string `json:"home_address"`
	Email           string `json:"email"`
	Gender          string   `json:"gender"`
	Telegram        string `json:"telegram"`
	Phone           string `json:"phone"`
	EmergencyName   string `json:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status"`
	Type            string   `json:"type"`
	Major           string `json:"major"`
	Graduate        string `json:"graduate"`
	Status          string `json:"status"`
	Category        string `json:"category"`
	Class           string `json:"class"`
}

func EntityToResponse(menteeEntity entity.MenteeEntity) response {
	return response{
		MenteeID:        menteeEntity.MenteeID,
		ClassID:         menteeEntity.ClassID,
		Class:           menteeEntity.Class,
		Name:            menteeEntity.Name,
		Address:         menteeEntity.Address,
		HomeAddress:     menteeEntity.HomeAddress,
		Email:           menteeEntity.Email,
		Telegram:        menteeEntity.Telegram,
		Phone:           menteeEntity.Phone,
		EmergencyName:   menteeEntity.EmergencyName,
		EmergencyPhone:  menteeEntity.EmergencyPhone,
		EmergencyStatus: menteeEntity.EmergencyStatus,
		Major:           menteeEntity.Major,
		Graduate:        menteeEntity.Graduate,
		Status:          menteeEntity.Status,
		Category:        menteeEntity.Category,
	}
}
