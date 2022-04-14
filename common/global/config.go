package global

import (
	"blog/common/model"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var Config *model.Config

func ParseConfig(path string) error {

	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return errors.New("配置文件不存在")
		} else {
			// Config file was found but another error was produced
			return err
		}
	}

	var conf model.Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Error("解析配置文件失败: %v", err)
	}
	Config = &conf
	return nil
}

func Init() {
	InitDb()
	InitLogger()
}
