/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	httpPort int = 8080
)

// metricsCmd represents the metrics command
var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Push metrics to prometheus",
	Long: `This command starts a http server having sigle endpoint /metrics.
These metrics can be scraped by prometheus.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("metrics called")
	},
}

func init() {
	rootCmd.AddCommand(metricsCmd)

	metricsCmd.Flags().IntVarP(&httpPort, "p", "port", 8080, "http port of the metric server")
}
