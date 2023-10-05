package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"restArchitecture/mikail/App/utils"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type RedisConfig struct {
	RedisAddr string `yaml:"redisAddr"`
	Password  string `yaml:"password"`
	DB        int    `yaml:"db"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Server ServerConfig   `yaml:"Server"`
	Mysql  DatabaseConfig `yaml:"Mysql"`
	Redis  RedisConfig    `yaml:"Redis"`
}

func StartConfig(env string) (*Config, error) {
	configName := utils.GetConfigEnv(env)
	file, err := os.ReadFile(configName)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
