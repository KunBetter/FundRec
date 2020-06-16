package store

type Persistence interface {
	Insert(data interface{})
	Update(data interface{}, where interface{})
	List(line interface{}, data []interface{})
	GetById(id interface{}, data interface{})
	Delete(id interface{}, data interface{})
}
