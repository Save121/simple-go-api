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
//TODO Add more database parameters
type DatabaseConfig struct {
	Port int `yaml:"port"`
}
func New()(*Settings, error) {
	var s Settings
	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}