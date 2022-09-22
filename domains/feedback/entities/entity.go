package feedbackentities

import (
	"mime/multipart"
	"time"
)

type FeedBackEntity struct {
	FeedBackID uint
	MenteeID   uint
	MentorID   uint
	MentorName string
	Status     string
	CreatedAt  time.Time
	Desc       string
	FileName   string
	FileData   multipart.File
	FileSize   int64
	Url        string
}
