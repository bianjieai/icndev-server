package initialization

import (
	"fmt"
	"github.com/bianjieai/icndev-server/configs"
	"github.com/bianjieai/icndev-server/internal/app/constant"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)

func NewMysqlDB(cfg configs.Mysql) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&time_zone=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
		encodeTimeZone(cfg.TimeZone),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("start mysql client failed, db:%s, err:%s", cfg.Database, err.Error())
	}
	return db
}

func encodeTimeZone(timezone string) string {
	if timezone == "" {
		timezone = constant.DefaultTimezone
	}

	return url.QueryEscape(fmt.Sprintf("'%s'", timezone))
}
