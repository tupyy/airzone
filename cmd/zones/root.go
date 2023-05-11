package zones

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
)

// rootZonesCmd represents the get command
var RootCmd = &cobra.Command{
	Use:   "zones",
	Short: "Control all the zones all together.",
	Run: func(c *cobra.Command, args []string) {
		hvac, err := hvac.GetData(common.Host, common.SystemID, common.AllZones)
		if err != nil {
			fmt.Printf("%+v", err)
			return
		}
		j, err := json.Marshal(hvac)
		if err != nil {
			fmt.Printf("%+v", err)
			return
		}
		fmt.Printf("%s", string(j))
	},
}
