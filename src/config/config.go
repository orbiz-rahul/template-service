package config

import (
	"encoding/json"
	"os"
)

type config struct {
	Postgres DBConfig `json:"PostgreSQL"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func GetTemplateConfig() (*config, error) {
	obj := new(config)

	filepath := "template_service_cfg.json"
	data, err := os.ReadFile(filepath)
	if err != nil {
		return obj, err
	}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return obj, err
	}
	return obj, nil

}
