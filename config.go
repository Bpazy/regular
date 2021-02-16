package regular

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

// configuration 总配置
type configuration struct {
	Telegram *telegramConfig `mapstructure:"telegram"`
	Proxy    *proxyConfig    `mapstructure:"proxy"`
}

// telegramConfig telegram 相关配置
type telegramConfig struct {
	Token string `mapstructure:"token"`
}

// proxyConfig 代理相关配置
type proxyConfig struct {
	Addr string `mapstructure:"addr"`
}

// check 校验配置文件必填项
func (c *configuration) check() {

}

func InitConfig() *configuration {
	viper.SetConfigName(".regular")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			createDefaultConfigFile()
		} else {
			log.Fatalf("保存配置文件失败: %+v", err)
		}
	}

	var c configuration
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("读取配置文件失败: %+v", err)
	}

	c.check()

	return &c
}

func createDefaultConfigFile() {
	viper.Set("telegram", telegramConfig{})
	viper.Set("proxy", proxyConfig{})

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatalf("初始化配置文件失败: %+v", err)
	}
	userHomeDir, _ := os.UserHomeDir()
	fmt.Println("请填写配置文件: " + filepath.Join(userHomeDir, ".regular.yaml"))
	os.Exit(0)
}
