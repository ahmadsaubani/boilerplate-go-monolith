package Config

import "time"

type AppConfig struct {
	Appname     string    `yaml:"name"`
	Debug       bool      `yaml:"debug"`
	Port        string    `yaml:"port"`
	Service     string    `yaml:"service"`
	Certificate string    `yaml:"certificate"`
	Pem_key     string    `yaml:"pem_key"`
	Host        string    `yaml:"host"`
	Api         ApiConfig `yaml:"api"`
}

type ApiConfig struct {
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}
