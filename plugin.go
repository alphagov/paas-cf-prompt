package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"flag"
	"fmt"
	"github.com/alphagov/paas-cf-prompt/prompt"
	"os"
	"strings"
)

type Plugin struct {
}

func (p *Plugin) Run(cliConnection plugin.CliConnection, args []string) {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	loggedOutPrompt := ""
	fs.StringVar(&loggedOutPrompt, "logged-out", "", "The prompt template to use when not logged in")
	err := fs.Parse(args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	org, err := cliConnection.GetCurrentOrg()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	space, err := cliConnection.GetCurrentSpace()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	api, err := cliConnection.ApiEndpoint()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	user, err := cliConnection.Username()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	version, err := cliConnection.ApiVersion()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	isLoggedIn, err := cliConnection.IsLoggedIn()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Join all args after the first together in to one string with spaces
	// in case the caller have given an unquoted string with spaces
	promptTemplate := strings.Join(fs.Args(), " ")
	if !isLoggedIn {
		promptTemplate = loggedOutPrompt
	}
	fmt.Println(prompt.Format(
		promptTemplate,
		prompt.PlaceholderValues{
			Org:     org.Name,
			Space:   space.Name,
			API:     api,
			User:    user,
			Version: version,
		},
	))
}

func (p *Plugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Prompt",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 26,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "prompt",
				HelpText: "Format a string for your terminal prompt (for example 'org: %o, space: %s')",
				UsageDetails: plugin.Usage{
					Usage: prompt.PROMPT_USAGE,
					Options: map[string]string{
						"logged-out": "logged out format string",
					},
				},
			},
		},
	}
}
