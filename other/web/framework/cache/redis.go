package cache

type RedisDriver struct {
}

func (r RedisDriver) Get(key string) interface{} {
	panic("implement me")
}

func (r RedisDriver) Set(key string, value interface{}) {
	panic("implement me")
}

func (r RedisDriver) Delete(key string) {
	panic("implement me")
}

func (r RedisDriver) Clear(key string) {
	panic("implement me")
}

func (r RedisDriver) Has(key string) {
	panic("implement me")
}
