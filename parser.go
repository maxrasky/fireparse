package fireparse

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	Header  Header  `yaml:"header"`
	Content Content `yaml:"content"`
}

func Parse(data string) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
