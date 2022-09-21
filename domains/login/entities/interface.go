package userentity

type ILoginRepo interface {
	GetByEmail(mentor MentorEntity) (MentorEntity, error)
}

type ILoginUsecase interface {
	Login(mentor MentorEntity) (string, error)
}
