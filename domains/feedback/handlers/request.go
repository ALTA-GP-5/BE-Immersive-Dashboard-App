package handlerfeedback

import (
	entity "immersive/domains/feedback/entities"
)

type request struct {
	MenteeID uint   `json:"mentee_id" form:"mentee_id" validate:"required"`
	Status   string `json:"status" form:"status" validate:"required"`
	Desc     string `json:"desc" form:"desc" validate:"required"`
}

func requestToEntity(feedbackRequest request) entity.FeedBackEntity {
	return entity.FeedBackEntity{
		MenteeID: feedbackRequest.MenteeID,
		Status:   feedbackRequest.Status,
		Desc:     feedbackRequest.Desc,
	}
}
