package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Cfg *Config

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	fmt.Println(".env files loaded")
	Cfg := &Config{
		ProjectName:      os.Getenv("PROJECT_NAME"),
		CompanyName:      os.Getenv("COMPANY_NAME"),
		ApplicationName:  os.Getenv("APPLICATION_NAME"),
		Environment:      os.Getenv("ENVIRONMENT"),
		MaxTableCapacity: GetIntEnv("MAX_TABLE_CAPACITY", 4),
	}
	return Cfg, nil
}

func GetIntEnv(envName string, defaultValue int) int {
	envValueStr := os.Getenv(envName)
	if envValueStr == "" {
		return defaultValue
	}
	envValue, err := strconv.Atoi(envValueStr)
	if err != nil {
		return defaultValue
	}
	return envValue
}
