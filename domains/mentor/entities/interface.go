package mentorentity

type IMentorRepo interface {
	Insert(mentor MentorEntity) error
	GetAll(mentor MentorEntity) ([]MentorEntity, error)
	GetById(mentor MentorEntity) (MentorEntity, error)
	Update(mentor MentorEntity) error
	Delete(mentor MentorEntity) error
}

type IMentorUsecase interface {
	Create(mentor MentorEntity) error
	ReadAll(mentor MentorEntity) ([]MentorEntity, error)
	ReadById(mentor MentorEntity) (MentorEntity, error)
	Update(mentor MentorEntity) error
	Delete(mentor MentorEntity) error
}
