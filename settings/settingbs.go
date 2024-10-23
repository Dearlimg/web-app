package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//func Init() (err error) {
//	viper.SetConfigFile("config.yaml")
//	//viper.SetConfigType("yaml")
//	viper.AddConfigPath("/settings/config.yaml")
//	err = viper.ReadInConfig()
//	if err != nil {
//		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
//		return
//	}
//	viper.WatchConfig()
//	viper.OnConfigChange(func(e fsnotify.Event) {
//		fmt.Println("配置文件修改了:", e.Name)
//	})
//	return
//}

func Init() (err error) {
	// 先设置配置文件的名称
	viper.SetConfigName("config")      // 不需要扩展名
	viper.SetConfigType("yaml")        // 确保设置配置类型
	viper.AddConfigPath("./settings/") // 当前目录
	//viper.AddConfigPath("/settings/") // 添加设置文件所在的目录，确保这是一个目录

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}

	// 开始监视配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件修改了:", e.Name)
	})
	return
}
