package cache

type cacheAdapterInterface interface {
	get(key string)
	set(key string, value interface{})
}
