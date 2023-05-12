/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package zone

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/cmd/common"
	"github.com/tupyy/airzone/internal/hvac"
)

// turnCmd represents the turn command
var turnCmd = &cobra.Command{
	Use:           "turn zoneID (on|off)",
	SilenceErrors: true,
	Short:         "Start or stop hvac for a zone",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("verb expected: either \"on\" or \"off\"")
		}

		zoneID, err := GetZoneID(context.TODO(), args[0])
		if err != nil {
			return err
		}

		action := args[1]
		resp := hvac.Hvac{}

		switch action {
		case "on":
			resp, err = hvac.Start(context.TODO(), common.Host, common.SystemID, zoneID, true)
		case "off":
			resp, err = hvac.Start(context.TODO(), common.Host, common.SystemID, zoneID, false)
		default:
			return fmt.Errorf("Expected either \"on\" or \"off\". Found %q", action)
		}

		if err != nil {
			return err
		}
		fmt.Printf("%+v", resp)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(turnCmd)

}
