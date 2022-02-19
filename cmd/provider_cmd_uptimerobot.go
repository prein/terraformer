package cmd

import (
	uptimerobot_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/uptimerobot"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdUptimeRobotImporter(options ImportOptions) *cobra.Command {
	var apiKey string
	cmd := &cobra.Command{
		Use:   "uptimerobot",
		Short: "Import current state to Terraform configuration from Uptime Robot",
		Long:  "Import current state to Terraform configuration from Uptime Robot",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newUptimeRobotProvider()
			err := Import(provider, options, []string{apiKey})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newUptimeRobotProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "user,team", "")
	cmd.PersistentFlags().StringVarP(&apiKey, "api-key", "", "", "Your Uptime Robot API key or env var UPTIMEROBOT_API_KEY")
	return cmd
}

func newUptimeRobotProvider() terraformutils.ProviderGenerator {
	return &uptimerobot_terraforming.UptimeRobotProvider{}
}
