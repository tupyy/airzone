package zones

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:           "set",
	Short:         "Set temperature or mode for all zones",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("temperature or mode arguments expected")
		}

		parameter := args[0]
		value := args[1]

		if parameter != "temperature" && parameter != "mode" {
			return fmt.Errorf("parameter to set %q unknown", parameter)
		}

		if parameter == "mode" {
			var m hvac.Mode
			switch value {
			case "cooling":
				m = hvac.CoollingMode
			case "heating":
				m = hvac.HeatingMode
			case "ventilation":
				m = hvac.VentilationMode
			case "dehumidification":
				m = hvac.Dehumidification
			default:
				return fmt.Errorf("unknown mode: %q", value)
			}

			resp, err := hvac.SetMode(context.TODO(), common.Host, common.SystemID, common.AllZones, m)
			if err != nil {
				return err
			}

			j, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			fmt.Printf("%s", string(j))
			return nil
		}

		temperature, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid temperature value: %q", value)
		}

		resp, err := hvac.SetTemperature(context.TODO(), common.Host, common.SystemID, common.AllZones, temperature)
		if err != nil {
			return err
		}

		j, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		fmt.Printf("%s", string(j))

		return nil
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
