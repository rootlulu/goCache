package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"cache/internal/app"

	reflogApp "reflog.com/reflog/pkg/app"
)

var config string

var rootCmd = &cobra.Command{
	Use:   "cache",
	Short: "the mysql cache",
	Long:  "the mysql cache",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debugf("cache start, cmd: %v, config: %v, args: %v", cmd, config, args)
		app.Run(cmd, config)
	},
	PreRun: Registerfunc,
}

func Registerfunc(cmd *cobra.Command, args []string) {
	reflogApp.RegisterFunc("GetUserCache", app.GetUserCache)
}

func main() {
	logrus.SetReportCaller(true)
	// go run cmd/server/main.go -c xxx.yaml
	rootCmd.Flags().StringVarP(&config, "config", "c", "", "the config file.")
	rootCmd.MarkFlagRequired("config")

	if err := rootCmd.Execute(); err != nil {
		panic(err.Error())
	}
}

func PrintArgs(args ...string) (string, error) {
	return reflogApp.PrintArgs(args...)
}
