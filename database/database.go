/**
 * @Author: huangw1
 * @Date: 2019/7/25 11:04
 */

package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	Dialect        string
	Url            string
	MaxIdle        int
	MaxActive      int
	EnableLogModel bool
	Models         []interface{}
}

var db *gorm.DB

func OpenDB(conf *DBConfig) *gorm.DB {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return fmt.Sprintf("t_%s", defaultTableName)
	}
	var err error
	db, err = gorm.Open(conf.Dialect, conf.Url)
	if err != nil {
		logrus.Fatalf("opens database failed: %s", err.Error())
	}
	db.SingularTable(true)

	maxIdle := 10
	if conf.MaxIdle > 0 {
		maxIdle = conf.MaxIdle
	}
	maxActive := 10
	if conf.MaxIdle > 0 {
		maxActive = conf.MaxIdle
	}
	db.LogMode(conf.EnableLogModel)
	db.DB().SetMaxIdleConns(maxIdle)
	db.DB().SetMaxOpenConns(maxActive)

	if len(conf.Models) > 0 {
		if err = db.AutoMigrate(conf.Models).Error; err != nil {
			logrus.Errorf("auto migrate tables failed: %s", err.Error())
		}
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		if err := db.Close(); err != nil {
			logrus.Errorf("closes database failed: %s", err.Error())
		}
	}
}

func Tx(db *gorm.DB, fun func(db *gorm.DB) error) (err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return
	}
	defer func() {
		if r := recover(); err != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()
	err = fun(db)
	return err
}
