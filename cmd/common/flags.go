package common

import (
	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/internal/presenter"
)

const (
	AllZones = 0
)

var (
	SystemID int = 1
	Host     string
)

func OutputPresenter(cmd *cobra.Command) func(v interface{}) (string, error) {
	output := cmd.Flag("output")
	if output == nil {
		return presenter.Json
	}

	if output.Value.String() == "table" {
		return presenter.Table
	}
	return presenter.Json
}
