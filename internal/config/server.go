package config

import "fmt"

type Server struct {
	Host    string `yaml:"host" env-default:"localhost"`
	Port    int    `yaml:"port" env-default:"8080"`
	Timeout int    `yaml:"timeout_in_s" env-default:"10"`
}

func (s Server) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
