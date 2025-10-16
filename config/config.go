package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type dataBase struct {
	DBHost     string `yaml:"dbHost"`
	DBPort     string `yaml:"dbPort"`
	DBName     string `yaml:"dbName"`
	DBUser     string `yaml:"dbUser"`
	DBPassword string `yaml:"dbPassword"`
}

type host struct {
	HostPort string `yaml:"hostPort"`
}

type Config struct {
	DataBase dataBase `yaml:"dataBase"`
	Host     host     `yaml:"host"`
}

func NewConfig() (*Config, error) {
	var config Config
	paths := []string{
		"config/config.yaml",    // если запускаем из корня
		"../config/config.yaml", // если запускаем из cmd/
		"./config.yaml",         // если файл рядом
	}

	var configFile *os.File
	var err error

	for _, path := range paths {
		configFile, err = os.Open(path)
		if err == nil {
			fmt.Println("✅ Нашел config.yaml по пути:", path)
			break
		}
	}

	if configFile == nil {
		return nil, fmt.Errorf("не могу найти config.yaml ни в одном месте")
	}

	defer configFile.Close()

	configBytes, err := io.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configBytes, &config)

	if err != nil {
		return nil, fmt.Errorf("error unmarshal yaml config: %w", err)
	}

	return &config, nil
}

func (config *Config) GetDsn() string {
	return "postgres://" + config.DataBase.DBUser + ":" + config.DataBase.DBPassword + "@" + config.DataBase.DBHost +
		":" + config.DataBase.DBPort + "/" + config.DataBase.DBName + "?sslmode=disable"
}
