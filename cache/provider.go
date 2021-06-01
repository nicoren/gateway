
package cache

type cacheAdapterProviderInterface interface {
	getCacheAdapter() *struct
}

type cacheAdapterProvider struct{
}

func (p cacheAdapterProvider) getCacheAdapter() *struct {
	switch c.adapter {
		case "redis":
			data := struct {
				adapterConfig cacheConfigAdapter
				cacheConfig   redisConfig
			}{}
		default:
			data := struct {
				adapterConfig cacheConfigAdapter
				cacheConfig   filesystemConfig
			}{}
	}
	return &data
}