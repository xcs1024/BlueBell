package mysql

import (
	"bluebell/domain"
	"bluebell/internal/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// Init 初始化mysql
func Init(cfg *setting.MySQLConfig) (err error) {
	//gorm的默认日志是只打印错误和慢Sql
	var mysqlLogger logger.Interface
	//要显示的日志等级
	mysqlLogger = logger.Default.LogMode(logger.Info)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})

	if err != nil {
		log.Fatalln("mysql connect failed", err)
		return err
	}
	//连接成功
	domain.DB = db
	log.Printf("mysql connect success")
	return
}
func Close() {
	db, _ := domain.DB.DB()
	_ = db.Close()
}
