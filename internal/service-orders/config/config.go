package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type SagaConfig struct {
	Retries   int `yaml:"retries"`
	TimeoutMs int `yaml:"timeout-ms"`
}

type Config struct {
	Service struct {
		Name string `yaml:"name"`
	} `yaml:"service"`
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
	Server struct {
		URL string `yaml:"url"`
	} `yaml:"server"`
	Metrics struct {
		PrometheusURL string `yaml:"prometheus-url"`
		JaegerURL     string `yaml:"jaeger-url"`
	} `yaml:"metrics"`
	Clients struct {
		Marketplace struct {
			URL     string `yaml:"url"`
			Timeout int64  `yaml:"timeout-ms"`
		} `yaml:"marketplace"`
		Warehouse struct {
			URL     string `yaml:"url"`
			Timeout int64  `yaml:"timeout-ms"`
		}
	} `yaml:"clients"`
	Saga SagaConfig `yaml:"saga"`
}

func New(filename string) (*Config, error) {
	var cfg Config

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config, err: <%v>", err)
	}

	if err = yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshall, err: <%v>", err)
	}
	return &cfg, nil
}
