package configs

import (
	"fmt"
	"os"

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
		ProjectName:     os.Getenv("PROJECT_NAME"),
		CompanyName:     os.Getenv("COMPANY_NAME"),
		ApplicationName: os.Getenv("APPLICATION_NAME"),
		Environment:     os.Getenv("ENVIRONMENT"),
	}
	return Cfg, nil
}
