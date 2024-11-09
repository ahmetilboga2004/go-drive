package interfaces

type IBaseRepository[T any] interface {
	GetAll() ([]*T, error)
	GetByID(id uint) (*T, error)
	Create(entity *T) error
	Update(id uint, entity *T) error
	Delete(id uint) error
}
