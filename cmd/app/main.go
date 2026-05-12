// Package main starts the Lab 2 Go Tooling application.

package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/Ruslann00/lab1-tooling/internal"
)

type Config struct {
	App struct {
		Name        string `mapstructure:"name"`
		Environment string `mapstructure:"environment"`
	} `mapstructure:"app"`

	Calculator struct {
		A int `mapstructure:"a"`
		B int `mapstructure:"b"`
	} `mapstructure:"calculator"`
}

func loadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return cfg
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Printf("failed to sync logger: %v\n", err)
		}
	}()

	cfg := loadConfig()

	a := cfg.Calculator.A
	b := cfg.Calculator.B

	sum := internal.Add(a, b)
	difference := internal.Subtract(a, b)

	result, err := internal.Divide(a, b)
	if err != nil {
		logger.Fatal("division failed", zap.Error(err))
	}

	logger.Info(
		"application started",
		zap.String("name", cfg.App.Name),
		zap.String("environment", cfg.App.Environment),
	)

	logger.Info(
		"calculation completed",
		zap.Int("a", a),
		zap.Int("b", b),
		zap.Int("sum", sum),
		zap.Int("difference", difference),
		zap.Int("division", result),
	)
}
