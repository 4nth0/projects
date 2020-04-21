package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type command struct {
	fs *flag.FlagSet
	fn func(args []string) error
}

var Version string
var ConfigPath string = ".projects/projects.yaml"

type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	commands := map[string]command{
		"list": listCmd(),
		"ls":   listCmd(),
	}

	fs := flag.NewFlagSet("projects", flag.ExitOnError)
	fs.Parse(os.Args[1:])
	args := fs.Args()

	if cmd, ok := commands[args[0]]; !ok {
		log.Fatalf("Unknown command: %s", args[0])
	} else if err := cmd.fn(args[1:]); err != nil {
		// help()
		log.Print(err)
	}
}
