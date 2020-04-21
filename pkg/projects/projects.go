package projects

import (
	"log"
	"os/exec"
)

type Project struct {
	Name        string   `yaml:"name,omitempty"`
	Description string   `yaml:"description,omitempty"`
	Path        string   `yaml:"path,omitempty"`
	Tags        []string `yaml:"tags,omitempty"`
}

func (p Project) Open(command string) {
	cmd := exec.Command(command, p.Path)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
