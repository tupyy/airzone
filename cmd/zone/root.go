package zone

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
)

// zoneCmd represents the zone command
var RootCmd = &cobra.Command{
	Use:          "zone",
	Short:        "Control one zone only.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Name or ZoneID is required")
		}

		// check if it is a zoneID or name
		zoneName := ""
		zoneID := -1
		zoneID, err := strconv.Atoi(strings.ToLower(args[0]))
		if err != nil {
			zoneName = strings.ToLower(args[0])
		}

		if zoneName != "" {
			names, err := hvac.GetZoneNames(common.Host, common.SystemID)
			if err != nil {
				return err
			}
			id, ok := names[zoneName]
			if !ok {
				return fmt.Errorf("Zone %q not found", zoneName)
			}
			zoneID = id
		}

		fmt.Println(zoneID)
		hvac, err := hvac.GetData(common.Host, common.SystemID, zoneID)
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
