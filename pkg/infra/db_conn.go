package infra

import (
	"fmt"

	"github.com/sakuradaman/go_simple_api/pkg/lib/config"
	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnector(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:@tcp(%s)/%s?parseTime=true&loc=Local", cfg.DBUser, cfg.DBAddress, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, xerrors.Errorf("db connection failed：%w", err)
	}
	return db, nil
}
