package settings

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var settingsFile []byte

type Settings struct {
	Port int `yaml:"port"`
	DB DatabaseConfig `yaml:"database"`
}
type DatabaseConfig struct {
	Port int `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Name string `yaml:"name"`
	Host string `yaml:"host"`

}
func New()(*Settings, error) {
	var s Settings
	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}