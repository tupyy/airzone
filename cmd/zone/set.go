package zone

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:           "set temperature zone value",
	Short:         "Set temperature for a particular zone",
	SilenceErrors: false,
	SilenceUsage:  false,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("temperature arguments expected")
		}

		parameter := args[0]
		zoneValue := args[1]
		value := args[2]

		if parameter != "temperature" {
			return fmt.Errorf("parameter to set %q unknown", parameter)
		}

		temperature, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid temperature value: %q", value)
		}

		zoneID, err := GetZoneID(context.TODO(), zoneValue)
		if err != nil {
			return err
		}

		resp, err := hvac.SetTemperature(context.TODO(), common.Host, common.SystemID, zoneID, temperature)
		if err != nil {
			return err
		}

		outputValue, err := common.OutputPresenter(cmd)(resp)
		if err != nil {
			return err
		}
		fmt.Printf("%s", outputValue)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
