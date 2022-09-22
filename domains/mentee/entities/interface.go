package menteeentites

type IMenteeRepo interface {
	Insert(menteeEntity MenteeEntity) error
	Update(menteeEntity MenteeEntity) error
	Delete(menteeEntity MenteeEntity) error
	FindAll(menteeEntity MenteeEntity) ([]MenteeEntity, error)
	Find(menteeEntity MenteeEntity) (MenteeEntity, error)
}

type IMenteeUseCase interface {
	Create(menteeEntity MenteeEntity) error
	Update(menteeEntity MenteeEntity) error
	Delete(menteeEntity MenteeEntity) error
	GetAll(menteeEntity MenteeEntity) ([]MenteeEntity, error)
	Get(menteeEntity MenteeEntity) (MenteeEntity, error)
}
