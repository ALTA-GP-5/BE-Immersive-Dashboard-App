package feedbackentities

import "time"

type FeedBackEntity struct {
	FeedBackID uint
	MenteeID   uint
	MentorID   uint
	MentorName string
	Status     string
	Date       time.Time
	Desc       string
	Url        string
}
