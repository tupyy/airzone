package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/cmd/zone"
	"github.com/tupyy/airzone/cmd/zones"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "airzone",
	Short: "Control your Airzone VAF",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(zones.RootCmd)
	RootCmd.AddCommand(zone.RootCmd)
	RootCmd.PersistentFlags().StringVarP(&common.Host, "host", "", "airzone:3000", "airzone host url. Example: 192.168.1.1:3000")
	RootCmd.PersistentFlags().IntVarP(&common.SystemID, "system-id", "", 1, "system id")
}
