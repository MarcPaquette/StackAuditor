package main

import (
	"flag"
	"fmt"
	"os"
	//	"strings"
	"code.cloudfoundry.org/cli/plugin"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type StackInventory struct {
	org   string
	space string
}

func main() {
	plugin.Start(new(StackInventory))
}

func (stackInventory *StackInventory) Run(cliConnection plugin.CliConnection, args []string) {
	//define and Parse args
	stackInventoryFlagSet := flag.NewFlagSet("stackinventory", flag.ExitOnError)
	org := stackInventoryFlagSet.String("org", "", "set org to inventory on")
	space := stackInventoryFlagSet.String("space", "", "set space to inventory on (requires -org)")

	err := stackInventoryFlagSet.Parse(args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *org == "" && *space != "" {
		fmt.Printf("Please set -org when using -space\n")
		os.Exit(1)
	}

	//Establish connection to capi
	apiEndpoint, err := cliConnection.ApiEndpoint()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfToken, err := cliConnection.AccessToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfconfig := &cfclient.Config{
		ApiAddress:        apiEndpoint,
		Token:             cfToken,
		SkipSslValidation: true, //TODO: make this configurable
	}

	client, err := cfclient.NewClient(cfconfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	apps, err := client.ListApps()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(apps)

	//Get GUID list of orgs filtering by org flag
	//Get GUID list of space from each org flitering by org/space flag
	//get GUID list of stacks
	//Get list of apps in each org & space with name and stack details
	//display list of apps (Org,Space, app, count of stack X, count of stack y, count of stack z) in tabular format

}

func (stackInventory *StackInventory) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "StackInventory",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "Stack Inventory",
				Alias:    "stackinventory",
				HelpText: "Reviews stacks in use by orgs and space. To obtain more information use --help",
				UsageDetails: plugin.Usage{
					Usage: "stackinventory - list stacks in use by org and space.\n   cf stackinventory [-org] [-space]",
					Options: map[string]string{
						"org":   "Specify the org to report",
						"space": "Specify the space to report (requires -org)",
					},
				},
			},
		},
	}
}
