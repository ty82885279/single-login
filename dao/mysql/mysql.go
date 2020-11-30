package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"single-login/setting"
)

var Db *gorm.DB

func Init(cfg *setting.MysqlConf) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	var s *sql.DB
	s, err = Db.DB()
	if err != nil {
		return err
	}
	return s.Ping()
}

func Close() {
	db, _ := Db.DB()
	_ = db.Close()
}
