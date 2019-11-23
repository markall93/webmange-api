package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mengjayxc/webmanage-api/internal/pkg/config"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models/db"
	"github.com/mengjayxc/webmanage-api/pkg/logger"
	"log"
	"os"
	"time"
	// import _ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB(config *config.Config) {
	var gdb *gorm.DB
	var err error
	if config.Gorm.DBType=="mysql"{
		config.Gorm.DSN=config.MySQL.DSN()
	} else if config.Gorm.DBType=="sqlite3"{
		config.Gorm.DSN=config.Sqlite3.DSN()
	}
	gdb, err = gorm.Open(config.Gorm.DBType, config.Gorm.DSN)
	if err != nil {
		panic(err)
	}
	gdb.SingularTable(true)
	if config.Gorm.Debug {
		gdb.LogMode(true)
		gdb.SetLogger(log.New(os.Stdout, "\r\n", 0))
	}
	gdb.DB().SetMaxIdleConns(config.Gorm.MaxIdleConns)
	gdb.DB().SetMaxOpenConns(config.Gorm.MaxOpenConns)
	gdb.DB().SetConnMaxLifetime(time.Duration(config.Gorm.MaxLifetime) * time.Second)
	db.DB=gdb
}

func CreateTables(tableStmts map[string]interface{}) {
	// Create all of the data tables
	for tableName := range tableStmts {
		has := db.DB.HasTable(tableName)
		if !has {
			table := tableStmts[tableName]
			_ = db.DB.CreateTable(table)
			logger.Debug("table:" + tableName + " " + "create successfully...")
		}
	}

}


