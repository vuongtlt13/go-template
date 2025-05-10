package config

import (
	"sync"
	"time"
	"yourapp/pkg/logger"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppMode  string       `env:"APP_MODE" envDefault:"production"`
	SqlDebug bool         `env:"SQL_DEBUG" envDefault:"false"`
	Cors     string       `env:"CORS" envDefault:"*"`
	Server   ServerConfig `envPrefix:"SERVER_"`
	Database DBConfig     `envPrefix:"DB_"`
	JWT      JWTConfig    `envPrefix:"JWT_"`
	Email    EmailConfig  `envPrefix:"EMAIL_"`
	Redis    RedisConfig  `envPrefix:"REDIS_"`
	Job      JobConfig    `envPrefix:"JOB_"`
}

func (c *Config) IsProduction() bool {
	return c.AppMode == "production"
}

type ServerConfig struct {
	UserPort  int `env:"USER_PORT" envDefault:"8080"`
	AdminPort int `env:"ADMIN_PORT" envDefault:"8081"`
}

type DBConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	User     string `env:"USER" envDefault:"postgres"`
	Password string `env:"PASSWORD" envDefault:"postgres"`
	Name     string `env:"NAME" envDefault:"postgres"`
}

type JWTConfig struct {
	Secret       string        `env:"SECRET,required"`
	ExpirePeriod time.Duration `env:"EXPIRE_PERIOD" envDefault:"24h"`
}

type EmailConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"1025"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	From     string `env:"FROM"`
}

type RedisConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"6379"`
	Password string `env:"PASSWORD" envDefault:""`
	DB       int    `env:"DB" envDefault:"0"`
}

type JobConfig struct {
	Concurrency int `env:"CONCURRENCY" envDefault:"5"`
}

var (
	instance *Config
	once     sync.Once
)

// GetConfig returns the singleton config instance
func GetConfig() *Config {
	var err error
	once.Do(func() {
		instance, err = loadConfig()
	})

	if err != nil {
		panic(err)
	}

	return instance
}

// loadConfig loads the configuration from file
func loadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		logger.GetLogger().Fatal("Failed to parse config: %v", err)
	}
	return cfg, nil
}
