package configuration

import "os"
import "gopkg.in/yaml.v2"

type Configuration struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func LoadConfig(path string) Configuration {
	file, err := os.Open(path)

	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	defer file.Close()

	var cfg Configuration
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		print(err.Error())
		os.Stderr.WriteString(err.Error())
	}

	return cfg
}
