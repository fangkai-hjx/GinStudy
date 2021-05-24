package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

//初始化MySQL连接池
func SetupDB() error {
	var err error
	DB, err = gorm.Open(DatabaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			DatabaseSetting.UserName,
			DatabaseSetting.PassWord,
			DatabaseSetting.Host,
			DatabaseSetting.DBName,
			DatabaseSetting.Charset,
			DatabaseSetting.ParseTime,
		))
	if err != nil {
		fmt.Println("gorm.Open err：", err)
		return err
	}
	if ServerSetting.RunMode == "debug" {
		DB.LogMode(true)
	}
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	DB.DB().SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	fmt.Println("MySQL pool init success")
	return nil
}
