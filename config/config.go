package config

import (
	"github.com/jinzhu/configor"
)

type AppConfig struct {
	Port uint `yaml:"port"`
}

type MySQLConfig struct {
	Host            string `yaml:"host"`
	Port            uint   `yaml:"port"`
	Db              string `yaml:"db"`
	UserName        string `yaml:"user_name"`
	Password        string `yaml:"password"`
	MaxIdleConn     int    `yaml:"max_idle_conn"`
	MaxOpenConn     int    `yaml:"max_open_conn"`
	ConnMaxLifeTime int    `yaml:"conn_max_life_time"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type WalletConfig struct {
	Appid     string `yaml:"appid"`
	SecretKey string `yaml:"secret_key"`
	Url       string `yaml:"url"`
	CallUrl   string `yaml:"call_url"`
}

type Config struct {
	App    AppConfig
	Mysql  MySQLConfig
	Redis  RedisConfig
	Wallet WalletConfig
}

var CONF *Config

func NewConfig(confPath string) (Config, error) {
	var config Config
	if confPath != "" {
		err := configor.Load(&config, confPath)
		if err != nil {
			return config, err
		}
	} else {
		err := configor.Load(&config, "config/config-example.yml")
		if err != nil {
			return config, err
		}
	}
	CONF = &config
	return config, nil
}
