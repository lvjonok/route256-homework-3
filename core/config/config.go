package config

type Config struct {
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
	Server struct {
		URL string `yaml:"url"`
	} `yaml:"server"`
}

func New(filename string) (*Config, error) {
	return nil, nil
}
