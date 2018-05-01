package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"io"
	"log"

	"github.com/urfave/cli"
)

var (
	columnSep = regexp.MustCompile(`\s+`)
	debug     = false
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
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Show debug logs",
			Destination: &debug,
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

func logD(format string, a ...interface{}) {
	if debug {
		if a == nil {
			log.Print(format)
		} else {
			log.Printf(format, a)
		}
	}
}

func logE(a ...interface{}) {
	if debug && a != nil {
		log.Fatal(a)
	}
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
	logD("execute start %v", cmd)
	var c *exec.Cmd
	if len(cmd) >= 2 {
		c = exec.Command(cmd[0], cmd[1:]...)
	} else {
		c = exec.Command(cmd[0])
	}

	if toStdin != "" {
		// I can't do anything if fails here.
		in, err := c.StdinPipe()
		if err != nil {
			logE(err)
		}

		go func() {
			_, err := io.WriteString(in, toStdin)
			if err != nil {
				logE(err)
			}
		}()
	}

	// Depending on the environment, fails here.
	s, err := c.Output()
	if err != nil {
		logE(err)
	}
	logD("execute finish %v", cmd)
	return string(s)
}

func removeHeader(in string) string {
	lines := strings.Split(in, "\n")

	// The measured value was 3.
	if len(lines) >= 3 {
		logD("removed a header.")
		return strings.Join(lines[1:], "\n")
	}

	// If there is no container or image,
	// remains header to avoiding error of peco.
	logD("remains a header.")
	return in
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
