package classmodel

import (
	entity "immersive/domains/feedback/entities"
	"time"

	"gorm.io/gorm"
)

type FeedBack struct {
	gorm.Model
	MentorID uint
	MenteeID uint
	Status   string
	Date     time.Time
	Desc     string
	Url      string

	Mentor Mentor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Mentee Mentee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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

type Mentee struct {
	gorm.Model
	ClassID  uint
	Name     string
	Gender   bool
	Status   string
	Category string
}

func EntityToModel(feedBackEntity entity.FeedBackEntity) FeedBack {
	return FeedBack{
		MentorID: feedBackEntity.MentorID,
		MenteeID: feedBackEntity.MenteeID,
		Status:   feedBackEntity.Status,
		Date:     feedBackEntity.Date,
		Desc:     feedBackEntity.Desc,
		Url:      feedBackEntity.Url,
	}
}

func ModelToEntity(feedBackModel FeedBack) entity.FeedBackEntity {
	return entity.FeedBackEntity{
		FeedBackID: feedBackModel.ID,
		MenteeID:   feedBackModel.MenteeID,
		MentorID:   feedBackModel.MentorID,
		MentorName: feedBackModel.Mentor.Fullname,
		Status:     feedBackModel.Status,
		Desc:       feedBackModel.Desc,
		Url:        feedBackModel.Url,
	}
}
