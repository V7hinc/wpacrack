package main

import (
	"github.com/urfave/cli/v2"
	"wpacrack/allCrack"
	"wpacrack/dictCrack"

	"log"
	"os"

	"wpacrack/logging"
)

var logger = logging.GetSugar()

func main() {
	app := cli.App{
		Name:  "wpa_crack",
		Usage: "",
		Commands: []*cli.Command{
			dictCrack.Command,
			allCrack.Command,
		},
		Flags: []cli.Flag{},
		Before: func(c *cli.Context) error {
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	logger.Infof("main over!")
}
