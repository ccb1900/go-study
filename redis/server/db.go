package server

import "errors"

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

type DBCollection struct {
	Collection []*DB
	Num        int
}

func NewDbList(num int) *DBCollection {
	dbList := make([]*DB, num)
	for i := 0; i < num; i++ {
		dbList[i] = NewDB(i)
	}

	return &DBCollection{
		Collection: dbList,
		Num:        num,
	}
}

func (d *DBCollection) Get(num int, k string) (*Object, error) {
	if v, exit := d.Collection[num].Store[k]; exit {
		return v, nil
	}

	return nil, errors.New("key not exists")
}

func (d *DBCollection) Set(dbNum int, k string, v *Object) {
	d.Collection[dbNum].Store[k] = v
}
