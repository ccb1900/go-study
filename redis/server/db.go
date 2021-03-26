package server

type DB struct {
	Store map[string]*Object
	Id    int
}

func NewDB(id int) *DB {
	return &DB{
		Store: make(map[string]*Object, 4096),
		Id:    id,
	}
}
