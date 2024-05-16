package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	Env              string `env:"ENV" env-default:"dev"`
	StoragePath      string `env:"STORAGE_PATH" env-required:"true"`
	HTTPServerConfig `yaml:"http_server"`
}

type HTTPServerConfig struct {
	Host             string        `env:"HOST" env-default:"localhost"`
	Port             string        `env:"PORT" env-default:"8080"`
	Timeout          time.Duration `env:"TIMEOUT" env-default:"4s"`
	KeepAliveTimeout time.Duration `env:"KEEP_ALIVE_TIMEOUT" env-default:"60s"`
}

/*
MustLoad загружает конфигурацию из файла и переменных окружения.

Практика именования функций вида MustSomething() - это распространенный шаблон именования в языках
программирования, таких как Go, и может применяться в других языках также.

В Go этот шаблон обычно используется для функций, которые выполняют некоторую
операцию, и в случае ошибки завершают программу или возвращают панику.

Например, если функция обязательно должна выполниться без ошибок иначе программа не сможет
продолжить работу, то её название может начинаться с "Must".
*/
func MustLoad() *Config {
	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config

	// Читаем конфигурацию из файла
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return &cfg
}
