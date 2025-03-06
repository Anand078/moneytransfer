package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"

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

func getProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filepath.Dir(filepath.Dir(b))) // 3 levels up to project root
	return basePath
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(getProjectRoot(), "configs"))

	if err := viper.ReadInConfig(); err != nil {
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
	filePath := filepath.Join(getProjectRoot(), "configs", "initial_balances.json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read initial balances file: %v", err)
	}

	var balances []InitialBalance
	if err := json.Unmarshal(data, &balances); err != nil {
		log.Fatalf("Failed to parse initial balances file: %v", err)
	}

	return balances
}
