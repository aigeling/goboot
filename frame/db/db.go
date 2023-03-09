package db

import (
	"fmt"
	"time"

	"github.com/aigeling/goboot/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbMap map[string]*gorm.DB

// 获得默认数据库
func GetDefaultDB() *gorm.DB {

	return GetDB(global.Conf.DefaultDB)
}

func GetDB(dbName string) *gorm.DB {
	db, ok := dbMap[dbName]
	if ok {
		return db
	} else {
		zap.L().Sugar().Errorf("db config not found:%s", dbName)
		return nil
	}
}

func InitDB() error {

	dbMap = make(map[string]*gorm.DB)
	debug := global.Conf.Debug
	for k, v := range global.Conf.DBConfigs {
		if !v.Enabled {
			continue
		}

		zap.L().Sugar().Infof("Initialize DB [%s] ...", k)
		db, err := CreateDB(&v, debug)
		if err != nil {

			zap.L().Sugar().Errorf("Init db[%s] err : %v", k, err)
			for _, e := range dbMap {
				_db, err := e.DB()
				if err != nil {
					continue
				}
				_db.Close()
			}
			return err

		}
		dbMap[k] = db
	}

	return nil
}

// 创建数据库
func CreateDB(o *global.DBConfig, debug bool) (db *gorm.DB, err error) {

	if o.DBType == "postgresql" {
		dnsFmt := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai"
		dsn := fmt.Sprintf(dnsFmt, o.Host, o.User, o.Password, o.DBName, o.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else if o.DBType == "mysql" {
		dnsFmt := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf(dnsFmt, o.User, o.Password, o.Host, o.Port, o.DBName)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		return
	}

	if debug {
		db = db.Debug()
	}

	sqlDB, _ := db.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(o.GetMaxOpenConns()) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(o.GetMaxIdleConns()) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于5，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxIdleTime(time.Duration(o.GetConnMaxIdleTime()) * time.Minute)

	return
}
