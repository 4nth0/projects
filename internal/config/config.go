package config

import (
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/4nth0/projects/pkg/projects"

	yaml "gopkg.in/yaml.v2"
)

// Config is the rout Config struct
type Config struct {
	path     string
	Command  string             `yaml:"command"`
	Projects []projects.Project `yaml:"projects"`
}

// LoadConfig load configuration yaml file content from the specified path
func LoadConfig(path string) *Config {
	t := Config{
		path: path,
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Err: ", err)
	}

	err = yaml.Unmarshal(data, &t)
	if err != nil {
		fmt.Println("error: %v", err)
	}

	return &t
}

func InitConfig(path string) *Config {
	cfg := Config{
		path: path,
	}

	return &cfg
}

func (c Config) Save() error {

	b, _ := yaml.Marshal(c)
	err := ioutil.WriteFile(c.path, b, 0644)
	if err != nil {
		fmt.Println("Err: ", err)
	}

	return nil
}

func GetConfigurationPathFromUserHome(relativePath string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir + "/" + relativePath, nil
}
