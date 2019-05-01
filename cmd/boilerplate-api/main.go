package main

import (
	"fmt"
	"os"

	"github.com/eventials/vlab-boilerplate-api/cmd/boilerplate-api/options/flags"
	"github.com/eventials/vlab-boilerplate-api/pkg/handlers"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

var (
	AppName    = "boilerplate-api"
	AppUsage   = "A cli tool to interact with vindi private api"
	AppVersion = "0.0.1"
	GitSummary = "none"
	GitBranch  = "none"
	GitMerge   = "0"
	CiBuild    = "0"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s summary=%s branch=%s merge=%s build=%s", c.App.Version, GitSummary, GitBranch, GitMerge, CiBuild)
	}

	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.EnableBashCompletion = true
	app.Version = AppVersion
	app.Commands = []cli.Command{}
	app.Flags = []cli.Flag{
		flags.Addr,
	}
	app.Action = func(c *cli.Context) error {
		addr := c.String("addr")
		return handlers.Serve(addr)
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
