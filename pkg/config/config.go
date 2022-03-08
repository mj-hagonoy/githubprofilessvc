package config

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Github struct {
		Username    string `yaml:"username"`
		AccessToken string `yaml:"access-token"`
		GetUserAPI  string `yaml:"get-user-api"`
	} `yaml:"github"`
	Cache struct {
		Type       string `yaml:"redis"`
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		DB         int    `yaml:"int"`
		ExpiryMins int64  `yaml:"expiry-mins"`
	} `yaml:"cache"`
}

func (c Config) ApiUrl() string {
	return fmt.Sprintf("http://%s:%s/api/v1", c.Host, c.Port)
}

func (c Config) RedisUrl() string {
	envValue := os.Getenv("REDIS_URL")
	if envValue == "" {
		return fmt.Sprintf("%s%s", c.Cache.Host, c.Cache.Port)
	}
	return envValue
}

var conf Config

func ParseConfig(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal([]byte(data), &conf); err != nil {
		return err
	}
	return nil
}

func GetConfig() Config {
	return conf
}

func NewConfig() *Config {
	return &conf
}
