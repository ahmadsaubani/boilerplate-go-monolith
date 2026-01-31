package Config

import (
	"time"
)

type DatabaseConfig struct {
	Name               string        `yaml:"name"`
	Username           string        `yaml:"username"`
	Password           string        `yaml:"password"`
	Port               string        `yaml:"port"`
	Engine             string        `yaml:"engine"`
	Host               string        `yaml:"host"`
	Maximum_connection int           `yaml:"maximum_connection"`
	MaximumIdleTime    time.Duration `yaml:"maximum_idle_time"`
	Usage              string        `yaml:"usage"`
	Connection         string        `yaml:"connection"`
}

type DatabasesConfig []DatabaseConfig
