package classentities

import "time"

type ClassEntity struct {
	ClassID       uint
	MentorID      uint
	Name          string
	Status        string
	StartDate     time.Time
	EndDate       time.Time
	MentorName    string
	GeneralSearch string
}
