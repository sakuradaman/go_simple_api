package infra

import (
	"fmt"

	"github.com/sakuradaman/go_simple_api/pkg/lib/config"
	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnector(cfg *config.Config) (*gorm.DB, error) {
	// DBに接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, xerrors.Errorf("db connection failed：%w", err)
	}
	return db, nil
}
