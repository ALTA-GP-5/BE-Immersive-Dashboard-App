package userentity

type IrepoUser interface {
	Store(userEntity UserEntity)
	Update(userEntity UserEntity) error
	Delete(userEntity UserEntity) error
	FindById(userEntity UserEntity) (UserEntity, error)
}

type IusecaseUser interface {
	Store(userEntity UserEntity)
	Update(userEntity UserEntity) error
	Delete(userEntity UserEntity) error
	GetById(userEntity UserEntity) (UserEntity, error)
}
