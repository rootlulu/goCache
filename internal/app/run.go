package app

import (
	"cache/internal/cache"
	"cache/internal/dao"
	"cache/pkg/utils"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var app = &App{}

type App struct {
	// TODO: add the close function.
	cmd *cobra.Command
	Cfg *viper.Viper
	Dao *dao.Dao

	Cache *cache.Cache
}

func initCfg(config string) {
	cfg, err := utils.Parse(config)
	if err != nil {
		logrus.Fatalln(err)
	}
	if key, ok := os.LookupEnv("cryptKey"); ok {
		cfg.Set("cryptKey", key)
	}
	app.Cfg = cfg
}

func initDao() {
	var err error
	app.Dao, err = dao.NewDao(app.Cfg)
	if err != nil {
		logrus.Fatalln(err)
	}
}

func initCache() {
	app.Cache = cache.NewCache(app.Dao)
}

// TODO: add the context
func Run(cmd *cobra.Command, config string) {
	app.cmd = cmd
	initCfg(config)
	initDao()
	initCache()
	initGrpcServer()
}

// RegisterFunc: logLevel
func SetLogLevel(args ...string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("GetCache need one arg")
	}
	var logLevel = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	if v, ok := logLevel[args[0]]; ok {
		logrus.SetLevel(v)
	} else {
		return "", fmt.Errorf("logLevel not found, must be one of debug, info, warn, error, fatal, panic")
	}
	return fmt.Sprintf("the logLevel set to %s", args[0]), nil
}
