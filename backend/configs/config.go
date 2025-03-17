package configs

import (
  "os"
  "log"
  "github.com/joho/godotenv"
)

type Config struct {
  Db DbConfig
}

type DbConfig struct {
  Dsn string
}

func LoadConfig() *Config {
  err := godotenv.Load()
  if err != nil {
    log.Println("[ ERROR ] can't load .env file, using default config")
  }
  return &Config{
    Db: DbConfig{
      Dsn: os.Getenv("DSN"),
    },
  }
}
