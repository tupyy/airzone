package zones

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
	"go.uber.org/zap"
)

// rootZonesCmd represents the get command
var RootCmd = &cobra.Command{
	Use:          "zones",
	Short:        "Control all the zones all together.",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		logger := common.SetupLogger(c.Flag("log-level").Value.String())
		defer logger.Sync()

		undo := zap.ReplaceGlobals(logger)
		defer undo()

		hvac, err := hvac.GetData(context.TODO(), common.Host, common.SystemID, common.AllZones)
		if err != nil {
			return err
		}

		value, err := common.OutputPresenter(c)(hvac)
		if err != nil {
			return err
		}
		fmt.Printf("%s", value)
		return nil
	},
}
