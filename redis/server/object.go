package server

type Object struct {
	Key   string
	Value string
}

func NewObject(k string, v string) *Object {
	return &Object{
		Key:   k,
		Value: v,
	}
}
