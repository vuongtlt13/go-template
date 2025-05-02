package database

import (
	"fmt"
	"sync"
	"time"
	"yourapp/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
	dbErr      error
)

type DBConfig struct {
	DSN             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func GetDatabase() *gorm.DB {
	once.Do(func() {
		appCfg := config.GetConfig()
		logMode := logger.Info
		if appCfg.SqlDebug {
			logMode = logger.Warn
		}

		// Connect to database
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			appCfg.Database.Host,
			appCfg.Database.Port,
			appCfg.Database.User,
			appCfg.Database.Password,
			appCfg.Database.Name,
		)

		dbInstance, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logMode),
		})
		if dbErr != nil {
			return
		}

		sqlDB, err := dbInstance.DB()
		if err != nil {
			dbErr = err
			return
		}

		sqlDB.SetMaxIdleConns(50)
		sqlDB.SetMaxOpenConns(200)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
	})

	if dbErr != nil {
		panic(dbErr)
	}
	return dbInstance
}
