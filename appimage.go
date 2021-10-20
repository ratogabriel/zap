package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/srevinsaju/zap/tui"
	"github.com/srevinsaju/zap/types"
	"github.com/urfave/cli/v2"
)

func installAppImageOptionsFromCLIContext(context *cli.Context) (types.InstallOptions, error) {
	executable := context.String("executable")
	appName := context.Args().First()

	from := context.String("from")
	if context.Bool("github") && from == "" {
		fmt.Printf("Installing from github requires the %s flag.\n", tui.Yellow("--from"))
		return types.InstallOptions{}, errors.New("github-from-flag-missing")
	}

	// use the repo name as appName
	if context.Bool("github") && appName == "" {
		fromSplit := strings.Split(from, "/")
		appName = fromSplit[len(fromSplit)-1]
	}

	if context.String("executable") == "" {
		logger.Debugf("Fallback executable name to appName, %s", context.Args().First())
		executable = appName
	}

	// get the second argument if provided, and if --from is not passed, to make the command line
	// interface more intuitive (originally suggested by @eadmaster at https://github.com/srevinsaju/zap/issues/31)
	fromFileAutomatic := context.Args().Get(1)
	if from == "" && fromFileAutomatic != "" {
		from = fromFileAutomatic
	}

	app := types.InstallOptions{
		Name:                   appName,
		From:                   from,
		Executable:             strings.Trim(executable, " "),
		FromGithub:             context.Bool("github"),
		RemovePreviousVersions: false,
		UpdateInplace:          context.Bool("update"),
		DoNotFilter:            context.Bool("no-filter"),
		Silent:                 context.Bool("silent"),
	}
	logger.Debug(app)
	return app, nil

}

func updateAppImageOptionsFromCLIContext(context *cli.Context) (types.Options, error) {
	executable := context.String("Executable")
	if context.String("Executable") == "" {
		executable = context.Args().First()
	}
	return types.Options{
		Name:              context.Args().First(),
		From:              context.String("from"),
		Executable:        executable,
		UseAppImageUpdate: context.Bool("with-au"),
		ForceRemove:       context.Bool("force-remove"),
	}, nil

}

func removeAppImageOptionsFromCLIContext(context *cli.Context) (types.RemoveOptions, error) {
	executable := context.String("Executable")
	if context.String("Executable") == "" {
		executable = context.Args().First()
	}
	return types.RemoveOptions{
		Executable: executable,
	}, nil

}
