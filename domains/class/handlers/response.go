package handlerclass

import (
	entity "immersive/domains/class/entities"
)

type response struct {
	Id        uint   `json:"id"`
	Mentor    string `json:"mentor"`
	Name      string `json:"name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"`
}

func EntityToResponse(classEntity entity.ClassEntity) response {
	return response{
		Id:        classEntity.ClassID,
		Mentor:    classEntity.MentorName,
		Name:      classEntity.Name,
		StartDate: classEntity.StartDate.Format("2006-01-02"),
		EndDate:   classEntity.EndDate.Format("2006-01-02"),
		Status:    classEntity.Status,
	}
}
