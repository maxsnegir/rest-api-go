package server

type Config struct {
	Port        string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	PostgresDsn string `toml:"postgres_dsn"`
	RedisDsn    string `toml:"redis_dsn"`
}

func NewConfig() *Config {
	return &Config{
		Port:     ":8000",
		LogLevel: "debug",
	}
}
