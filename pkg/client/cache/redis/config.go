package redis

// RedisConfig is an interface for redis configs.
type RedisConfig interface {
	Address() string
	Password() string
}
