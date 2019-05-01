package flags

import "github.com/urfave/cli"

var (
	Addr = cli.StringFlag{Name: "addr", Value: ":8080", Usage: "Addr to bind the web server", EnvVar: "VLAB_BOILERPLATE_ADDR"}
)
