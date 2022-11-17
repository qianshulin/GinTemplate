package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dev  bool   `yaml:"Dev"`
	Host string `yaml:"Host"`
	Auth struct {
		AccessSecret string `yaml:"AccessSecret"`
		AccessExpire int64  `yaml:"AccessExpire"`
	} `yaml:"Auth"`
	Mongo struct {
		DataSource string `yaml:"DataSource"`
		Database   string `yaml:"Database"`
	} `yaml:"Mongo"`
	ETHUrl string `yaml:"ETHUrl"`

	AssetURL string `yaml:"AssetURL"`
}

func Init(path string) (*Config, error) {
	c := &Config{}
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(content, c); err != nil {
		return nil, err
	}

	return c, nil
}
