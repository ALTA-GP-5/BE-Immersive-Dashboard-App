package classentities

type IClassRepo interface {
	Insert(classEntity ClassEntity) error
	Update(classEntity ClassEntity) error
	Delete(classEntity ClassEntity) error
	FindAll(classEntity ClassEntity) ([]ClassEntity, error)
	Find(classEntity ClassEntity) (ClassEntity, error)
}

type IClassUseCase interface {
	Create(classEntity ClassEntity) error
	Update(classEntity ClassEntity) error
	Delete(classEntity ClassEntity) error
	GetAll(classEntity ClassEntity) ([]ClassEntity, error)
	Get(classEntity ClassEntity) (ClassEntity, error)
}
