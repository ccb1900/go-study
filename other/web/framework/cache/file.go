package cache

type FileDriver struct {
}

func (f FileDriver) Get(key string) interface{} {
	panic("implement me")
}

func (f FileDriver) Set(key string, value interface{}) {
	panic("implement me")
}

func (f FileDriver) Delete(key string) {
	panic("implement me")
}

func (f FileDriver) Clear(key string) {
	panic("implement me")
}

func (f FileDriver) Has(key string) {
	panic("implement me")
}
