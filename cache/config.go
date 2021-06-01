package cache

type cacheConfigAdapter struct {
	adapter string
}

func (c cacheConfig) getCacheConfigStruct() *struct {
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
