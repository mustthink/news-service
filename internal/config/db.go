package config

import "fmt"

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"pass"`
	DBName   string `yaml:"db"`
}

func (d Database) Uri() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		d.Host, d.Port, d.User, d.DBName, d.Password)
}
