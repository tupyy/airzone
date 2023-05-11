package zones

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
)

// rootZonesCmd represents the get command
var RootCmd = &cobra.Command{
	Use:          "zones",
	Short:        "Control all the zones all together.",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		hvac, err := hvac.GetData(context.TODO(), common.Host, common.SystemID, common.AllZones)
		if err != nil {
			return err
		}
		j, err := json.Marshal(hvac)
		if err != nil {
			return err
		}
		fmt.Printf("%s", string(j))
		return nil
	},
}
