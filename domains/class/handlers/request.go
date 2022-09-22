package handlerclass

import (
	entity "immersive/domains/class/entities"
)

type request struct {
	MentorID  uint   `json:"mentor_id" form:"mentor_id" validate:"required"`
	Name      string `json:"name" form:"name" validate:"required"`
	Status    string `json:"status" form:"status" validate:"required"`
	StartDate string `json:"start_date" form:"start_date" validate:"required"`
	EndDate   string `json:"end_date" form:"end_date" validate:"required"`
}

func requestToEntity(classRequest request) entity.ClassEntity {
	return entity.ClassEntity{
		MentorID: classRequest.MentorID,
		Name:     classRequest.Name,
		Status:   classRequest.Status,
	}
}
