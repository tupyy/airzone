package zone

import (
	"fmt"

	"github.com/spf13/cobra"
)

// zoneCmd represents the zone command
var RootCmd = &cobra.Command{
	Use:   "zone",
	Short: "Control one zone only.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zone called")
	},
}
