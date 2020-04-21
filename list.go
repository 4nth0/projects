package main

import (
	"flag"
	"os"
	"strings"

	"github.com/4nth0/projects/internal/config"
	"github.com/4nth0/projects/pkg/projects"

	"github.com/manifoldco/promptui"
)

var detailsTemplate = `
--------- Project ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}
{{ "Path:" | faint }}	{{ .Path }}
{{ "Tags:" | faint }}	{{range $index, $element := .Tags}}{{if $index}}, {{end}}{{.}}{{end}}`

func listCmd() command {
	fs := flag.NewFlagSet("projects list", flag.ExitOnError)

	return command{fs, func(args []string) error {
		return listProjects()
	}}
}

func listProjects() error {
	var p []projects.Project

	configAbsolutePath, err := config.GetConfigurationPathFromUserHome(ConfigPath)
	if err != nil {
		return err
	}

	conf := config.LoadConfig(configAbsolutePath)
	tags := os.Args[2:]

	p = projects.FilterProjectsByTag(conf.Projects, tags)
	project, err := AskForProject(p)
	if err == nil {
		project.Open(conf.Command)
		return nil
	} else {
		return err
	}
}

func AskForProject(p []projects.Project) (*projects.Project, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000025B8 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "\U000025B8 {{ .Name | white | faint }}",
		Details:  detailsTemplate,
	}

	searcher := func(input string, index int) bool {
		project := p[index]
		name := strings.Replace(strings.ToLower(project.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Project to open",
		Items:     p,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return &p[index], nil
}
