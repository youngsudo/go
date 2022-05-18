package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() error {
	// viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 查找配置文件所在的路径
	viper.AddConfigPath("./settings/") // 是相对于整个工程目录的路径
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			fmt.Println("配置文件未找到")
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println("配置文件被找到，但产生了另外的错误")
		}
		return err
	}

	// 监听配置文件变化
	viper.WatchConfig()
	// 当配置文件发生变化时，会触发OnConfigChange方法
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	return nil
}
