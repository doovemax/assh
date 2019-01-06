package commands // import "github.com/doovemax/assh/pkg/commands"

import (
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"

	"github.com/doovemax/assh/pkg/config"
)

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	abspath, err := filepath.Abs(ex)
	if err != nil {
		log.Fatal(err)
	}
	config.SetASSHBinaryPath(abspath)
}

// Commands is the list of cli commands
var Commands = []cli.Command{
	{
		Name:        "connect",
		Usage:       "Connect to host SSH socket, used by ProxyCommand",
		Description: "Argument is a host.",
		Action:      cmdProxy,
		Hidden:      true,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "no-rewrite",
				Usage: "Do not automatically rewrite outdated configuration",
			},
			cli.IntFlag{
				Name:  "port, p",
				Usage: "SSH destination port",
			},
			cli.BoolFlag{
				Name:  "dry-run",
				Usage: "Only show how assh would connect but don't actually do it",
			},
		},
	},
	{
		Name:   "ping",
		Usage:  "Send packets to the SSH server and display statistics",
		Action: cmdPing,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port, p",
				Usage: "SSH destination port",
			},
			cli.UintFlag{
				Name:  "count, c",
				Usage: "Stop after sending 'count' packets",
			},
			cli.Float64Flag{
				Name:  "wait, i",
				Usage: "Wait 'wait' seconds between sending each packet",
				Value: 1,
			},
			cli.BoolFlag{
				Name:  "o",
				Usage: "Exit successfully after receiving one reply packet",
			},
			cli.Float64Flag{
				Name:  "waittime, W",
				Usage: "Time in seconds to wait for a reply for each packet sent",
				Value: 1,
			},
		},
	},
	/*
		{
			Name:        "info",
			Usage:       "Print the connection config for host",
			Description: "Argument is a host.",
			Action:      cmdInfo,
		},
	*/
	/*
		{
			Name:        "init",
			Usage:       "Configure SSH to use assh",
			Description: "Build a .ssh/config.advanced file based on .ssh/config and update .ssh/config to use assh as ProxyCommand.",
			Action:      cmdInit,
		},
	*/
	/*
		{
			Name:        "etc-hosts",
			Usage:       "Generate a /etc/hosts file with assh hosts",
			Description: "Build a .ssh/config.advanced file based on .ssh/config and update .ssh/config to use assh as ProxyCommand.",
			Action:      cmdEtcHosts,
		},
	*/
	{
		Name:   "info",
		Usage:  "Display system-wide information",
		Action: cmdInfo,
	},
	{
		Name:  "config",
		Usage: "Manage ssh and assh configuration",
		Subcommands: []cli.Command{
			{
				Name:  "build",
				Usage: "Build .ssh/config",
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "no-automatic-rewrite",
						Usage: "Disable automatic ~/.ssh/config file regeneration",
					},
					cli.BoolFlag{
						Name:  "expand, e",
						Usage: "Expand all fields",
					},
					cli.BoolFlag{
						Name:  "ignore-known-hosts",
						Usage: "Ignore known-hosts file",
					},
				},
				Action: cmdBuild,
			},
			{
				Name:   "json",
				Usage:  "Returns the JSON output",
				Action: cmdBuildJSON,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "expand, e",
						Usage: "Expand all fields",
					},
				},
			},
			{
				Name:   "list",
				Usage:  "List all hosts from assh config",
				Action: cmdList,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "expand, e",
						Usage: "Expand all fields",
					},
				},
			},
			{
				Name:   "graphviz",
				Usage:  "Generate a Graphviz graph of the hosts",
				Action: cmdGraphviz,
				Flags: []cli.Flag{
					cli.BoolFlag{
						Name:  "show-isolated-hosts",
						Usage: "Show isolated hosts",
					},
					cli.BoolFlag{
						Name:  "no-resolve-wildcard",
						Usage: "Do not resolve wildcards in Gateways",
					},
					cli.BoolFlag{
						Name:  "no-inheritance-links",
						Usage: "Do not show inheritance links",
					},
				},
			},
			{
				Name:   "search",
				Usage:  "Search entries by given search text",
				Action: cmdSearch,
			},
		},
	},
	{
		Name:  "sockets",
		Usage: "Manage control sockets",
		Subcommands: []cli.Command{
			{
				Name:   "list",
				Action: cmdCsList,
				Usage:  "List active control sockets",
			},
			{
				Name:   "flush",
				Action: cmdCsFlush,
				Usage:  "Close control sockets",
			},
			{
				Name:   "master",
				Action: cmdCsMaster,
				Usage:  "Open a master control socket",
			},
		},
	},
	// FIXME: tree
	{
		Name:   "wrapper",
		Usage:  "Initialize assh, then run ssh/scp/rsync...",
		Hidden: true,
		Subcommands: []cli.Command{
			{
				Name:   "ssh",
				Action: cmdWrapper,
				Usage:  "Wrap ssh",
				Flags:  config.SSHFlags,
			},
		},
	},
}
