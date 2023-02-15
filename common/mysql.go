package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/prometheus/common/log"
)

// MysqlConfig 创建结构体
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

// GetMysqlFromConsul 获取mysql的配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	//获取配置
	if err := config.Get(path...).Scan(mysqlConfig); err != nil {
		log.Error(err)
		return nil
	}
	return mysqlConfig
}
