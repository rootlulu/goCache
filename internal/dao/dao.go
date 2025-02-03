package dao

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"cache/internal/config"
	"cache/pkg/utils"
)

type Dao struct {
	Mysql *gorm.DB
	Redis string
	Mongo string
}

func NewDao(vp *viper.Viper) (*Dao, error) {
	var dao = &Dao{}
	mysqlConn, err := NewMysqlConn(vp)
	if err != nil {
		return nil, fmt.Errorf("mysql connected failed: %s", err)
	}
	dao.Mysql = mysqlConn
	return dao, nil
}

func NewMysqlConn(vp *viper.Viper) (*gorm.DB, error) {
	var cfg = config.Mysql{}
	if err := vp.UnmarshalKey("mysql", &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal 'mysql' config: %v", err)
	}

	var db *gorm.DB
	var err error
	var uri string
	fmt.Println(vp.GetString("cryptKey"))
	if cfg.URI != "" {
		if vp.IsSet("cryptKey") {
			// the encrypted.
			uri, err = utils.DeCrypt(cfg.URI, []byte(vp.GetString("cryptKey")))
		} else {
			uri = cfg.URI
		}
		if err != nil {
			return nil, err
		}
		db, err = gorm.Open(mysql.Open(uri), &gorm.Config{})
	} else {
		// TODO: parse the confgi map.
	}
	if err != nil {
		logrus.Fatalf("connected to mysql failed: %v", err)
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("get mysql db failed: %v", err)
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		logrus.Fatalf("connected to mysql failed: %v", err)
		return nil, err
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
	return db, nil
}
