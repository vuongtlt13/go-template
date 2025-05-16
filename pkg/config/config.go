package config

import (
	"sync"
	"time"
	"yourapp/pkg/logger"

	"github.com/caarlos0/env/v11"
	"github.com/spf13/viper"
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
	I18n     I18nConfig   `envPrefix:"I18N_"`
}

func (c *Config) IsProduction() bool {
	return c.AppMode == "production"
}

type ServerConfig struct {
	Port int    `env:"PORT" mapstructure:"port" envDefault:"8000"`
	Cors string `env:"CORS" mapstructure:"cors"`
	App  AppConfig
}

type AppConfig struct {
	Name         string        `env:"NAME" mapstructure:"name"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" mapstructure:"idle_timeout"`
	RateLimit    RateLimitConfig
}

type RateLimitConfig struct {
	Max        int           `env:"MAX" mapstructure:"max"`
	Expiration time.Duration `env:"EXPIRATION" mapstructure:"expiration"`
}

type DBConfig struct {
	Host     string `env:"HOST" mapstructure:"host" envDefault:"localhost"`
	Port     int    `env:"PORT" mapstructure:"port" envDefault:"5432"`
	User     string `env:"USER" mapstructure:"user" envDefault:"postgres"`
	Password string `env:"PASSWORD" mapstructure:"password" envDefault:"postgres"`
	Name     string `env:"NAME" mapstructure:"name" envDefault:"postgres"`
}

type JWTConfig struct {
	Secret       string        `env:"SECRET" mapstructure:"secret,required"`
	ExpirePeriod time.Duration `env:"EXPIRE_PERIOD" mapstructure:"expire_period" envDefault:"24h"`
}

type EmailConfig struct {
	Host     string `env:"HOST" mapstructure:"host" envDefault:"localhost"`
	Port     int    `env:"PORT" mapstructure:"port" envDefault:"1025"`
	Username string `env:"USERNAME" mapstructure:"username"`
	Password string `env:"PASSWORD" mapstructure:"password"`
	From     string `env:"FROM" mapstructure:"from"`
}

type RedisConfig struct {
	Host     string `env:"HOST" mapstructure:"host" envDefault:"localhost"`
	Port     int    `env:"PORT" mapstructure:"port" envDefault:"6379"`
	Password string `env:"PASSWORD" mapstructure:"password" envDefault:""`
	DB       int    `env:"DB" mapstructure:"db" envDefault:"0"`
}

type JobConfig struct {
	Concurrency int `env:"CONCURRENCY" mapstructure:"concurrency" envDefault:"5"`
}

type I18nConfig struct {
	DefaultLocale string `env:"DEFAULT_LOCALE" mapstructure:"default_locale" envDefault:"en"`
	BaseFolder    string `env:"BASE_FOLDER" mapstructure:"base_folder" envDefault:"i18n/locales"`
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

// loadConfig loads the configuration from system environment variables
func loadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		logger.GetLogger().Fatal("Failed to parse config: %v", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Get is an alias for GetConfig for backward compatibility
func Get() *Config {
	return GetConfig()
}
