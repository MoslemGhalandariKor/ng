package redis


import (
    "context"
    "github.com/redis/go-redis/v9"
    "log"
)

var (
    Client *redis.Client
    Ctx    = context.Background()
)

type RedisConfig struct {
    RedisHost                string `mapstructure:"redis_host"`
    RedisPort            string `mapstructure:"redis_port"`
    RedisPassword       string `mapstructure:"redis_password"`
	RedisDB   			int  `mapstructure:"redis_db"`
}
func Initialize(config RedisConfig) {
    Client = redis.NewClient(&redis.Options{
        Addr:     config.RedisHost + ":" + config.RedisPort,
        Password: config.RedisPassword,
        DB:       config.RedisDB,
    })

    _, err := Client.Ping(Ctx).Result()
    if err != nil {
        log.Fatalf("❌ Failed to connect to Redis: %v", err)
    }

    log.Println("✅ Connected to Redis")
}
