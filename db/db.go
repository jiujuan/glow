package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jiujuan/glow/config"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func InitDB() {
	mysqlConf := config.BlogConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@/%s?charset=%s", mysqlConf.Username, mysqlConf.Password, mysqlConf.DB, mysqlConf.Charset)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Errorf("open db error: %s", err.Error())
	}
	db.DB().SetMaxIdleConns(mysqlConf.MaxIdleConn)
	db.DB().SetMaxOpenConns(mysqlConf.MaxOpenConn)
	db.SingularTable(true)
	db.LogMode(true)

	if err = db.DB().Ping(); err != nil {
		log.Errorf("ping error: %s", err.Error())
	}
	DB = db
}
