package config

import "strings"

type Config struct {
	getEnv func(string) string
}

func NewConfig(getEnv func(string) string) *Config {
	return &Config{
		getEnv: getEnv,
	}
}

func (c *Config) GetEnv(key string) string {
	return c.getEnv(strings.ToUpper(key))
}

func (c *Config) GetEnvOrDefault(key, fallback string) string {
	if v := c.GetEnv(key); len(v) > 0 {
		return v
	}
	return fallback
}
