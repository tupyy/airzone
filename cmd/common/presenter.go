package common

import (
	"github.com/spf13/cobra"
	"github.com/tupyy/airzone/internal/presenter"
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
