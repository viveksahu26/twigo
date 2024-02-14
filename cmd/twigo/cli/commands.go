package cli

import (
	"github.com/spf13/cobra"
	"github.com/viveksahu26/twigo/cmd/twigo/cli/templates"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "twigo",
		Short:             "tool to send messages using SMS, Whatapps, Voice note.",
		DisableAutoGenTag: true,
		SilenceUsage:      true, // Don't show usage on errors
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	// ro.AddFlags(cmd)

	// templates.SetCustomUsageFunc(cmd)

	// Add sub-commands.

	// cmd.AddCommand(Chat())
	cmd.AddCommand(Send())
	templates.SetCustomUsageFunc(cmd)

	return cmd
}
