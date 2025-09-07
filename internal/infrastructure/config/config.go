package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App      App      `yaml:"app"`
		HTTP     HTTP     `yaml:"http"`
		Log      Log      `yaml:"log"`
		PG       PG       `yaml:"postgres"`
		Kafka    Kafka    `yaml:"kafka"`
		Redis    Redis    `yaml:"redis"`
	}

	App struct {
		Name    string `yaml:"name" env:"APP_NAME" env-default:"myapp"`
		Version string `yaml:"version" env:"APP_VERSION" env-default:"1.0.0"`
	}

	HTTP struct {
		Port        string        `yaml:"port" env:"HTTP_PORT" env-default:"8080"`
		ReadTimeout time.Duration `yaml:"read_timeout" env:"HTTP_READ_TIMEOUT" env-default:"5s"`
	}

	Log struct {
		Level string `yaml:"level" env:"LOG_LEVEL" env-default:"debug"`
	}

	PG struct {
		URL         string `yaml:"url" env:"PG_URL" env-default:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
		MaxPoolSize int    `yaml:"max_pool_size" env:"PG_MAX_POOL_SIZE" env-default:"10"`
	}

	Kafka struct {
		Brokers []string `yaml:"brokers" env:"KAFKA_BROKERS" env-separator:"," env-default:"localhost:9092"`
		Topic   string   `yaml:"topic" env:"KAFKA_TOPIC" env-default:"events"`
		GroupID string   `yaml:"group_id" env:"KAFKA_GROUP_ID" env-default:"mygroup"`
	}

	Redis struct {
		Addr     string        `yaml:"addr" env:"REDIS_ADDR" env-default:"localhost:6379"`
		Password string        `yaml:"password" env:"REDIS_PASSWORD" env-default:""`
		DB       int           `yaml:"db" env:"REDIS_DB" env-default:"0"`
		TTL      time.Duration `yaml:"ttl" env:"REDIS_TTL" env-default:"1h"`
	}
)

var (
	instance *Config
 	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config-local.yaml", instance); err != nil {
			panic(err)
		}
	})
	return instance
}