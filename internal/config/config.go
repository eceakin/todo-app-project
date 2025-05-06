package config

type Config struct {
	JWTSecret string
}

func NewConfig() *Config {
	return &Config{
		JWTSecret: "your_secret_key", // Gerçek projede bu değer env'den alınmalı
	}
}
