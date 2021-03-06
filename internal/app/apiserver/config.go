package apiserver

type Config struct {
	BindAddr  string `toml:"bind_addr"`
	LogLevel  string `toml:"log_level"`
	CacheAddr string `toml:"cache_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:  ":8080",
		LogLevel:  "debug",
		CacheAddr: ":11211",
	}
}
