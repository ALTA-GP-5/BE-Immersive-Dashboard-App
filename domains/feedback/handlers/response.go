package handlerfeedback

import (
	entity "immersive/domains/feedback/entities"
)

type response struct {
	FeedBackID uint   `json:"feedback_id"`
	Status     string `json:"status"`
	MentorName string `json:"mentor_name"`
	Desc       string `json:"desc"`
	Date       string `json:"date"`
	Url        string `json:"url"`
}

func EntityToResponse(feedBackEntity entity.FeedBackEntity) response {
	return response{
		FeedBackID: feedBackEntity.FeedBackID,
		Status:     feedBackEntity.Status,
		MentorName: feedBackEntity.MentorName,
		Desc:       feedBackEntity.Desc,
		Date:       feedBackEntity.CreatedAt.Format("2006-01-02"),
		Url:        feedBackEntity.Url,
	}
}
