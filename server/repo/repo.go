package repo

import "github.com/go-redis/redis"

type CountLogger interface {
	Fatalf(format string, a ...interface{})
	Warningf(foramt string, a ...interface{})
	Infof(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	V(int) bool
}

type CountRepoRedis struct {
	addr   string
	redis  *redis.Client
	logger CountLogger
}

func NewCountRepoRedis(addr string, logger CountLogger) *CountRepoRedis {
	red := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if err := red.Ping().Err(); err != nil {
		logger.Fatalf("failed to connect redis: %v", err)
	}
	logger.Infof("connected to redis")

	return &CountRepoRedis{
		addr:   addr,
		redis:  red,
		logger: logger,
	}
}

func (r *CountRepoRedis) Incr(key string) (int, error) {
	ret, err := r.redis.Incr(key).Result()
	return int(ret), err
}
