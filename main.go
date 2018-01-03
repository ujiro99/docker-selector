package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

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
		var ids string
		if image {
			ids = execDockerImage()
		} else if all {
			ids = execDockerPsAll()
		} else {
			ids = execDockerPs()
		}
		fmt.Print(ids)
		return nil
	}

	app.Run(os.Args)
}

func execDockerImage() string {
	var s string
	s = output([]string{"docker", "images"}, "")
	s = removeHeader(s)
	s = output([]string{"peco"}, s)
	return extractColumn(s, 2)
}

func execDockerPs() string {
	var s string
	s = output([]string{"docker", "ps"}, "")
	s = removeHeader(s)
	s = output([]string{"peco"}, s)
	return extractColumn(s, 0)
}

func execDockerPsAll() string {
	var s string
	s = output([]string{"docker", "ps", "-a"}, "")
	s = removeHeader(s)
	s = output([]string{"peco"}, s)
	return extractColumn(s, 0)
}

func output(cmd []string, toStdin string) string {
	var c *exec.Cmd
	if len(cmd) >= 2 {
		c = exec.Command(cmd[0], cmd[1:]...)
	} else {
		c = exec.Command(cmd[0])
	}
	if toStdin != "" {
		// I can't do anything if fails here.
		in, _ := c.StdinPipe()
		io.WriteString(in, toStdin)
		in.Close()
	}
	// Depending on the environment, fails here.
	s, err := c.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}

func removeHeader(out string) string {
	lines := strings.Split(out, "\n")
	// The measured value was 3.
	if len(lines) >= 3 {
		return strings.Join(lines[1:], "\n")
	}
	// If there is no container or image,
	// remains header to avoiding error of peco.
	return out
}

func extractColumn(out string, pos int) string {
	lines := strings.Split(out, "\n")
	ids := make([]string, len(lines))
	for i, l := range lines {
		columns := columnSep.Split(l, -1)
		if 0 <= pos && pos < len(columns) {
			ids[i] = columns[pos]
		}
	}
	return strings.Join(ids, " ")
}
