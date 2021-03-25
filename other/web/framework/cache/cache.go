package cache

type ICache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Delete(key string)
	Clear(key string)
	Has(key string)
}

func NewCache() ICache {
	return new(RedisDriver)
}
