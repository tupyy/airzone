package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	systemID int = 1
	zoneID   int = 0
	host     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hvac",
	Short: "Control your Airzone VAF",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "host", "", "airzone:3000", "airzone host url. Example: 192.168.1.1:3000")
	rootCmd.PersistentFlags().IntVarP(&systemID, "system-id", "", 1, "system id. default to 1")
	rootCmd.PersistentFlags().IntVarP(&zoneID, "zone-id", "", 0, "zone id. Defaults to 0 (all zones)")
}
