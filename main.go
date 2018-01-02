package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mattn/go-pipeline"
	"github.com/urfave/cli"
)

var (
	columnSep = regexp.MustCompile(`\s+`)
)

func main() {
	var all bool
	var image bool
	app := cli.NewApp()
	app.Name = "docker-selector"
	app.Usage = "docker container selector."

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "all, a",
			Usage:       "Show all containers",
			Destination: &all,
		},
		cli.BoolFlag{
			Name:        "image, i",
			Usage:       "Show images",
			Destination: &image,
		},
	}

	app.Action = func(c *cli.Context) error {
		var id string
		if image {
			id = execDockerImage()
		} else if all {
			id = execDockerPecoAll()
		} else {
			id = execDockerPeco()
		}
		fmt.Print(id)
		return nil
	}

	app.Run(os.Args)
}

func execDockerImage() string {
	out, err := pipeline.Output(
		[]string{"docker", "images"},
		[]string{"peco"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return extractColumn(out, 2)
}

func execDockerPeco() string {
	out, err := pipeline.Output(
		[]string{"docker", "ps"},
		[]string{"peco"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return extractColumn(out, 0)
}

func execDockerPecoAll() string {
	out, err := pipeline.Output(
		[]string{"docker", "ps", "-a"},
		[]string{"peco"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return extractColumn(out, 0)
}

func extractColumn(dockerOut []byte, pos int) string {
	lines := strings.Split(string(dockerOut), "\n")
	ids := make([]string, len(lines))
	for i, l := range lines {
		columns := columnSep.Split(l, -1)
		if 0 <= pos && pos < len(columns) {
			ids[i] = columns[pos]
		}
	}
	return strings.Join(ids, " ")
}
