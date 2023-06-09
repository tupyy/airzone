package zone

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
	"go.uber.org/zap"
)

// zoneCmd represents the zone command
var RootCmd = &cobra.Command{
	Use:          "zone",
	Short:        "Control one zone only.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := common.SetupLogger(cmd.Flag("log-level").Value.String())
		defer logger.Sync()

		undo := zap.ReplaceGlobals(logger)
		defer undo()

		if len(args) == 0 {
			return errors.New("Name or ZoneID is required")
		}

		zoneID, err := GetZoneID(context.TODO(), args[0])
		if err != nil {
			return err
		}

		hvac, err := hvac.GetData(context.TODO(), common.Host, common.SystemID, zoneID)
		if err != nil {
			return err
		}

		value, err := common.OutputPresenter(cmd)(hvac)
		if err != nil {
			return err
		}
		fmt.Printf("%s", value)

		return nil
	},
}

func GetZoneID(ctx context.Context, arg string) (int, error) {
	// check if it is a zoneID or name
	zoneName := ""
	zoneID := -1
	zoneID, err := strconv.Atoi(strings.ToLower(arg))
	if err != nil {
		zoneName = strings.ToLower(arg)
	}

	if zoneName != "" {
		names, err := hvac.GetZoneNames(ctx, common.Host, common.SystemID)
		if err != nil {
			return 0, err
		}
		id, ok := names[zoneName]
		if !ok {
			return 0, fmt.Errorf("Zone %q not found", zoneName)
		}
		zoneID = id
	}

	return zoneID, nil
}
