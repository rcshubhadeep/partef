package main

import (
	"fmt"
	"os"

	"github.com/rcshubhadeep/partef/jsonfile"

	"github.com/urfave/cli"
)

func main() {

	var concReq int
	var debug bool
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "conc, c",
			Value:       10,
			Usage:       "Number of concurent requests",
			Destination: &concReq,
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Debug mode",
			Destination: &debug,
		},
	}

	app.Name = "partef"
	app.Usage = "Load testing of any REST API."
	app.Version = "0.1"
	app.UsageText = "partef [global options] <filename.json>"

	app.Action = func(c *cli.Context) error {
		fileName := "api.json"
		if c.NArg() > 0 {
			fileName = c.Args()[0]
		}
		if !jsonfile.FileExists(fileName) {
			fmt.Println("File does not exists")
			return nil
		}
		err := jsonfile.ProcessAPIDesc(fileName, concReq)
		if err != nil {
			if debug {
				fmt.Println(err.Error())
			}
		}
		return nil
	}

	app.Run(os.Args)
}
