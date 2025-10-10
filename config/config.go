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
	pathToYamlFile := "../../config/config.yaml"
	configFile, err := os.Open(pathToYamlFile)
	if err != nil {
		return nil, fmt.Errorf("error decode config file: %w", err)
	}

	defer configFile.Close() //Закрытие файла confige после отработки текущей функции

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
	return "postgres://" + config.DataBase.DBUser + ":" + config.DataBase.DBPassword + "@" + config.DataBase.DBHost + ":" + config.DataBase.DBPassword + "/" + config.DataBase.DBName + "?sslmode=disable"
}
