package config

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	Name            string        `yaml:"name"             env:"APP_NAME"             env-default:"altura-api"`
	Env             string        `yaml:"env"              env:"APP_ENV"              env-default:"development"`
	Port            int           `yaml:"port"             env:"APP_PORT"             env-default:"8080"`
	Debug           bool          `yaml:"debug"            env:"APP_DEBUG"            env-default:"false"`
	ReadTimeout     time.Duration `yaml:"read_timeout"     env:"APP_READ_TIMEOUT"     env-default:"10s"`
	WriteTimeout    time.Duration `yaml:"write_timeout"    env:"APP_WRITE_TIMEOUT"    env-default:"10s"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"APP_SHUTDOWN_TIMEOUT" env-default:"30s"`
}

type DatabaseConfig struct {
	Host            string        `yaml:"host"              env:"DB_HOST"              env-default:"localhost"`
	Port            int           `yaml:"port"              env:"DB_PORT"              env-default:"5432"`
	User            string        `yaml:"user"              env:"DB_USER"              env-default:"postgres"`
	Password        string        `yaml:"password"          env:"DB_PASSWORD"`
	Name            string        `yaml:"name"              env:"DB_NAME"              env-default:"altura"`
	SSLMode         string        `yaml:"sslmode"           env:"DB_SSLMODE"           env-default:"disable"`
	MaxOpenConns    int           `yaml:"max_open_conns"    env:"DB_MAX_OPEN_CONNS"    env-default:"25"`
	MaxIdleConns    int           `yaml:"max_idle_conns"    env:"DB_MAX_IDLE_CONNS"    env-default:"5"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" env:"DB_CONN_MAX_LIFETIME" env-default:"5m"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" env:"DB_CONN_MAX_IDLE_TIME" env-default:"1m"`
	ConnectTimeout  time.Duration `yaml:"connect_timeout"   env:"DB_CONNECT_TIMEOUT"   env-default:"10s"`
}

type CacheConfig struct {
	Host         string        `yaml:"host"          env:"REDIS_HOST"          env-default:"localhost"`
	Port         int           `yaml:"port"          env:"REDIS_PORT"          env-default:"6379"`
	Password     string        `yaml:"password"      env:"REDIS_PASSWORD"`
	DB           int           `yaml:"db"            env:"REDIS_DB"            env-default:"0"`
	PoolSize     int           `yaml:"pool_size"     env:"REDIS_POOL_SIZE"     env-default:"10"`
	DialTimeout  time.Duration `yaml:"dial_timeout"  env:"REDIS_DIAL_TIMEOUT"  env-default:"5s"`
	ReadTimeout  time.Duration `yaml:"read_timeout"  env:"REDIS_READ_TIMEOUT"  env-default:"3s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"REDIS_WRITE_TIMEOUT" env-default:"3s"`
}

type LoggerConfig struct {
	Level string `yaml:"level" env:"LOG_LEVEL" env-default:"info"`
}

type Config struct {
	App      AppConfig      `yaml:"app"`
	Database DatabaseConfig `yaml:"database"`
	Cache    CacheConfig    `yaml:"cache"`
	Logger   LoggerConfig   `yaml:"logger"`
}

func (a *AppConfig) IsProduction() bool {
	return a.Env == "production"
}
func (a *AppConfig) IsDevelopment() bool {
	return a.Env == "development"
}
func (a *AppConfig) IsStaging() bool {
	return a.Env == "staging"
}

var (
	validEnvs = map[string]bool{
		"development": true,
		"staging":     true,
		"production":  true,
	}
	validSSLModes = map[string]bool{
		"disable": true, "require": true, "verify-ca": true, "verify-full": true,
	}
	validLogLevels = map[string]bool{
		"debug": true, "info": true, "warn": true, "error": true, "fatal": true,
	}
)

func (cfg *Config) Validate() error {
	var errs []error
	if cfg.App.Name == "" {
		errs = append(errs, errors.New("app.name is required"))
	}
	if !validEnvs[cfg.App.Env] {
		errs = append(errs, fmt.Errorf("app.env %q must be one of: development, staging, production", cfg.App.Env))
	}
	if cfg.App.Port < 1 || cfg.App.Port > 65535 {
		errs = append(errs, fmt.Errorf("app.port %d is out of range (1-65535)", cfg.App.Port))
	}
	if cfg.App.ReadTimeout <= 0 {
		errs = append(errs, errors.New("app.read_timeout must be positive"))
	}
	if cfg.App.WriteTimeout <= 0 {
		errs = append(errs, errors.New("app.write_timeout must be positive"))
	}
	if cfg.App.ShutdownTimeout <= 0 {
		errs = append(errs, errors.New("app.shutdown_timeout must be positive"))
	}

	if cfg.Database.Host == "" {
		errs = append(errs, errors.New("database.host is required"))
	}
	if cfg.Database.Port < 1 || cfg.Database.Port > 65535 {
		errs = append(errs, fmt.Errorf("database.port %d is out of range (1-65535)", cfg.Database.Port))
	}
	if cfg.Database.User == "" {
		errs = append(errs, errors.New("database.user is required"))
	}
	if cfg.Database.Password == "" && cfg.App.IsProduction() {
		errs = append(errs, errors.New("database.password is required in production"))
	}
	if cfg.Database.Name == "" {
		errs = append(errs, errors.New("database.name is required"))
	}
	if !validSSLModes[cfg.Database.SSLMode] {
		errs = append(errs, fmt.Errorf("database.sslmode %q must be one of: disable, require, verify-ca, verify-full", cfg.Database.SSLMode))
	}
	if cfg.Database.MaxOpenConns < 1 {
		errs = append(errs, errors.New("database.max_open_conns must be at least 1"))
	}
	if cfg.Database.MaxIdleConns < 0 {
		errs = append(errs, errors.New("database.max_idle_conns must be non-negative"))
	}
	if cfg.Database.MaxIdleConns > cfg.Database.MaxOpenConns {
		errs = append(errs, errors.New("database.max_idle_conns must not exceed max_open_conns"))
	}

	if cfg.Cache.Host == "" {
		errs = append(errs, errors.New("cache.host is required"))
	}
	if cfg.Cache.Port < 1 || cfg.Cache.Port > 65535 {
		errs = append(errs, fmt.Errorf("cache.port %d is out of range (1-65535)", cfg.Cache.Port))
	}
	if cfg.Cache.PoolSize < 1 {
		errs = append(errs, errors.New("cache.pool_size must be at least 1"))
	}

	if !validLogLevels[cfg.Logger.Level] {
		errs = append(errs, fmt.Errorf("logger.level %q must be one of: debug, info, warn, error, fatal", cfg.Logger.Level))
	}

	return errors.Join(errs...)
}

func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	var cfg Config
	configPath := fmt.Sprintf("configs/config.%s.yml", env)
	if _, err := os.Stat(configPath); err == nil {
		if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
			return nil, fmt.Errorf("read config.%s.yml: %w", env, err)
		}
	} else {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			return nil, fmt.Errorf("read env: %w", err)
		}
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &cfg, nil
}

func (d *DatabaseConfig) DSN() string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(d.User, d.Password),
		Host:   fmt.Sprintf("%s:%d", d.Host, d.Port),
		Path:   d.Name,
	}
	q := u.Query()
	q.Set("sslmode", d.SSLMode)
	q.Set("connect_timeout", strconv.Itoa(int(d.ConnectTimeout.Seconds())))
	u.RawQuery = q.Encode()
	return u.String()
}

func (c *CacheConfig) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
