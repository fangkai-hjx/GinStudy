package global

import (
	"GinStudy/pkg/setting"
	"fmt"
	"time"
)

//服务器配置
type ServerSettings struct {
	RunMode     string
	HttpHost    string
	HttpPort    string
	ReadTimeout time.Duration
}

//数据库配置
type DatabaseSettings struct {
	DBType       string
	UserName     string
	PassWord     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
type RedisSettings struct {
	Addr     string
	Password string
}

//定义全局变量
var (
	ServerSetting   *ServerSettings
	DatabaseSetting *DatabaseSettings
	RedisSetting    *RedisSettings
)

//读取配置到全局变量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		fmt.Println("setting.NewSetting is err：", err)
		return err
	}
	err = s.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis", &RedisSetting)
	if err != nil {
		return err
	}
	return nil
}
