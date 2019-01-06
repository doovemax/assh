package commands

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/bugsnag/osext"
	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/doovemax/assh/pkg/config"
	"github.com/doovemax/assh/pkg/utils"
	"github.com/doovemax/assh/pkg/version"
)

func cmdInfo(c *cli.Context) error {
	conf, err := config.Open(c.GlobalString("config"))
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	fmt.Printf("Debug mode (client): %v\n", os.Getenv("ASSH_DEBUG") == "1")
	cliPath, err := osext.Executable()
	if err != nil {
		return err
	}
	fmt.Printf("CLI Path: %s\n", cliPath)
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("Version: %s (%s)\n", version.VERSION, version.GITCOMMIT)
	fmt.Println("")
	fmt.Printf("RC files:\n")
	homeDir := utils.GetHomeDir()
	for _, filename := range conf.IncludedFiles() {
		relativeFilename := strings.Replace(filename, homeDir, "~", -1)
		fmt.Printf("- %s\n", relativeFilename)
	}
	fmt.Println("")
	fmt.Println("Statistics:")
	fmt.Printf("- %d hosts\n", len(conf.Hosts))
	fmt.Printf("- %d templates\n", len(conf.Templates))
	fmt.Printf("- %d included files\n", len(conf.IncludedFiles()))
	// FIXME: print info about connections/running processes
	// FIXME: print info about current config file version

	return nil
}
