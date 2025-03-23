package config

import (
	"log"
	"os"

	"github.com/Netflix/go-env"
)

type Config struct {
	Mongo struct {
		Url string `env:"MONGO_URL,default=mongodb://mongo:FjLIDhMUpmlMeaQtBkxUqoOjlHZuxPRW@hopper.proxy.rlwy.net:15301"`
	}

	Nats struct {
		Url string `env:"NATS_URL,default=nats://localhost:4222"`
	}

	Patient struct {
		Port int `env:"PORT,default=8081"`
	}

	JWT struct {
		Secret string `env:"JWT_SECRET,default=secret"`
	}

	Auth struct {
		Port int `env:"PORT,default=8080"`
	}

	Environment string `env:"ENVIRONMENT,default=development"`

	Log struct {
		Enable bool `env:"LOG_ENABLE,default=true"`
	}

	Version struct {
		Name string `env:"API_VERSION,default=unset"`
		Hash string `env:"GIT_COMMIT_HASH,default=unset"`
	}

	Redis struct {
		URL string `env:"REDIS_URL,default=redis://localhost:6379/11"`
	}

	Extras env.EnvSet
}

func LoadConfig() *Config {
	var cfg Config
	es, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Remaining environment variables.
	cfg.Extras = es
	return &cfg
}
