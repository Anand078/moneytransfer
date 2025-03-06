package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type InitialBalance struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type Config struct {
	Server          ServerConfig     `mapstructure:"server"`
	InitialBalances []InitialBalance `mapstructure:"initial_balances"`
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadConfig(); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	conf.InitialBalances = loadInitialBalancesFile()

	return conf
}

func loadInitialBalancesFile() []InitialBalance {
	viper.SetConfigFile("configs/initial_balances.json")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read initial balances: %v", err)
	}

	var balances []InitialBalance
	if err := viper.Unmarshal(&balances); err != nil {
		log.Fatalf("Failed to parse initial balances: %v", err)
	}
	return balances
}
