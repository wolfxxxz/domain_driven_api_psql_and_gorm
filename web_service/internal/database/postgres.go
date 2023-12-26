package database

import (
	"context"
	"fmt"
	"strconv"
	"time"
	"web_service/internal/apperrors"
	"web_service/internal/config"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresDB interface {
	SetupDatabase(ctx context.Context, conf *config.Config, log *zap.Logger) (*gorm.DB, error)
}

type postgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB() PostgresDB {
	return &postgresDB{}
}

func (p *postgresDB) SetupDatabase(ctx context.Context, conf *config.Config, log *zap.Logger) (*gorm.DB, error) {
	if conf.Postgres.DBName == "" {
		appErr := apperrors.SetupDatabaseErr.AppendMessage("config DBName is empty")
		log.Sugar().Error(appErr)
		return nil, appErr
	}

	dsn := fmt.Sprintf("%v://%v:%v/%v?sslmode=%v&user=%v&password=%v&dbname=%v",
		conf.Postgres.SqlType, conf.Postgres.SqlHost, conf.Postgres.SqlPort, conf.Postgres.SqlType,
		conf.Postgres.SqlMode, conf.Postgres.UserName, conf.Postgres.Password, conf.Postgres.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		appErr := apperrors.SetupDatabaseErr.AppendMessage(err)
		log.Sugar().Error(appErr)
		return nil, appErr
	}

	sqlDB, err := db.DB()
	if err != nil {
		appErr := apperrors.SetupDatabaseErr.AppendMessage(err)
		log.Sugar().Error(appErr)
		return nil, appErr
	}

	tNum, err := strconv.Atoi(conf.Postgres.TimeoutQuery)
	if err != nil {
		appErr := apperrors.SetupDatabaseErr.AppendMessage(err)
		log.Sugar().Error(appErr)
		return nil, appErr
	}

	dsnWithoutPassword := fmt.Sprintf("%v://%v:%v/%v?sslmode=%v&user=%v&password=[great secret]&dbname=%v&TimeZone=%s",
		conf.Postgres.SqlType, conf.Postgres.SqlHost, conf.Postgres.SqlPort, conf.Postgres.SqlType,
		conf.Postgres.SqlMode, conf.Postgres.UserName, conf.Postgres.DBName, conf.Postgres.TimeZone,
	)
	log.Sugar().Infof("Trying to connect to Postgres.\n %s", dsnWithoutPassword)
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(tNum))
	defer cancel()

	err = sqlDB.PingContext(ctx)
	if err != nil {
		appErr := apperrors.SetupDatabaseErr.AppendMessage(fmt.Sprintf("PingErr %v", err))
		log.Sugar().Error(appErr)
		return nil, appErr
	}

	log.Info("DB Postgres has been connected, DB.Ping success ")
	return db, nil
}
