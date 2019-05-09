package main // import "src.techknowlogick.com/drone-nfpm"

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "nfpm plugin"
	app.Usage = "nfpm plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "config file",
			EnvVar: "PLUGIN_CONFIG,CONFIG",
		},
		cli.StringFlag{
			Name:   "target",
			Usage:  "where to save the generated package",
			EnvVar: "PLUGIN_TARGET,TARGET",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Config: c.String("config"),
		Target:	c.String("target"),
	}

	return plugin.Exec()
}