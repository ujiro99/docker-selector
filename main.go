package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/mattn/go-pipeline"
	"log"
	"strings"
)

func main() {
	var all bool;
	app := cli.NewApp()
	app.Name = "docker-selector"
	app.Usage = "docker container selector."

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "all, a",
			Usage:       "Show all containers",
			Destination: &all,
		},
	}

	app.Action = func(c *cli.Context) error {
		var ps string
		if all {
			ps = execDockerPecoAll()
		} else {
			ps = execDockerPeco()

		}
		id := extractId(ps)
		fmt.Print(id)
		return nil
	}

	app.Run(os.Args)
}

func execDockerPeco() string {
	out, err := pipeline.Output(
		[]string{"docker", "ps", "--format", "{{.ID}}\t{{.Image}}\t{{.Status}}\t{{.Command}}\t{{.RunningFor}}"},
		[]string{"peco"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func execDockerPecoAll() string {
	out, err := pipeline.Output(
		[]string{"docker", "ps", "-a", "--format", "{{.ID}}\t{{.Image}}\t{{.Status}}\t{{.Command}}\t{{.RunningFor}}"},
		[]string{"peco"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func extractId(processes string) string {
	pos := strings.Index(processes, "\t")
	return processes[0:pos]
}
