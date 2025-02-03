package main

// import (
// 	"github.com/sirupsen/logrus"
// 	"github.com/spf13/cobra"

// 	"reflog/pkg/app"
// )

// var logLevel string

// var rootCmd = &cobra.Command{
// 	Use:   "cache",
// 	Short: "the mysql cache",
// 	Long:  "the mysql cache",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		logrus.Debugf("cache start, cmd: %v, logLevel: %v, args: %v", cmd, logLevel, args)
// 		app.Run(cmd, logLevel)
// 	},
// 	PreRun: Registerfunc,
// }

// func Registerfunc(cmd *cobra.Command, args []string) {
// 	// ...
// }

// func main() {
// 	logrus.SetReportCaller(true)
// 	// go run cmd/server/main.go -c xxx.yaml
// 	rootCmd.Flags().StringVarP(&logLevel, "logLevel", "l", "", "the log level")
// 	rootCmd.MarkFlagRequired("logLevel")

// 	if err := rootCmd.Execute(); err != nil {
// 		panic(err.Error())
// 	}
// }
