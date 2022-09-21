package mentorhandler

import entity "immersive/domains/mentor/entities"

type response struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Team     string `json:"team"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

func EntityToResponse(mentor entity.MentorEntity) response {
	return response{
		ID:       mentor.MentorID,
		FullName: mentor.FullName,
		Email:    mentor.Email,
		Team:     mentor.Team,
		Role:     mentor.Role,
		Status:   mentor.Status,
	}
}
