package server

type Config struct {
	Port        string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		Port:     ":8000",
		LogLevel: "debug",
	}
}
