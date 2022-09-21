package handlerclass

import (
	entity "immersive/domains/class/entities"
	"time"
)

type response struct {
	Id        uint      `json:"id"`
	Mentor    string    `json:"mentor"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"`
}

func EntityToResponse(classEntity entity.ClassEntity) response {
	return response{
		Id:        classEntity.ClassID,
		Mentor:    classEntity.MentorName,
		Name:      classEntity.Name,
		StartDate: classEntity.StartDate,
		EndDate:   classEntity.EndDate,
		Status:    classEntity.Status,
	}
}
