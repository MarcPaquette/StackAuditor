package main

import (
	//	"flag"
	"fmt"
	//	"os"
	//	"strings"

	"code.cloudfoundry.org/cli/plugin"
)

type StackAudit struct {
	org   string
	space string
}

func main() {
	plugin.Start(new(StackAudit))
}

func (stackAuditor *StackAudit) Run(cliConnection plugin.CliConnection, args []string) {
	fmt.Printf("hello\n")
}

func (stackAuditor *StackAudit) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "StackAudit",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "Stack Auditor",
				Alias:    "stackaudit",
				HelpText: "Reviews stacks in use by orgs and space. To obtain more information use --help",
				UsageDetails: plugin.Usage{
					Usage: "stackaudit - list stacks in use by org and space.\n   cf stackaudit [-org] [-space]",
					Options: map[string]string{
						"org":   "Specify the org to report",
						"space": "Specify the space to report (requires -org)",
					},
				},
			},
		},
	}
}
