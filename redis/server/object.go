package server

type Object struct {
	Key   string
	Value string
	DBNum int
}

func NewObject(dbNum int, k string, v string) *Object {
	return &Object{
		Key:   k,
		Value: v,
		DBNum: dbNum,
	}
}
