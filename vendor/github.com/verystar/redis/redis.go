package redis

import (
	"github.com/go-redis/redis"
	"log"
	"strings"
)

var (
	redisList map[string]*redis.Client
	errs      []string
)

type Config struct {
	Server   string
	Password string
	DB       int
}

func Client(name ... string) *redis.Client {
	key := "default"
	if name != nil {
		key = name[0]
	}

	client, ok := redisList[key]
	if !ok {
		log.Fatalf("[redis] the redis client `%s` is not configured", key)
	}

	return client
}

func Connect(configs map[string]Config) {
	defer func() {
		if len(errs) > 0 {
			log.Fatal("[redis] " + strings.Join(errs, "\n"))
		}
	}()

	redisList = make(map[string]*redis.Client)
	for name, conf := range configs {
		r := newRedis(&conf)
		log.Println("[redis] connect:" + conf.Server)

		_, err := r.Ping().Result()
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}

		client := newRedis(&conf)

		if r, ok := redisList[name]; ok {
			redisList[name] = client
			r.Close()
		} else {
			redisList[name] = client
		}
	}
}

// 创建 redis pool
func newRedis(conf *Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Server,
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
	})
	return client
}
