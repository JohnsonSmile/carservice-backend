package config

import (
	"carservice/infra/logger"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	serverConfig *ServerConfig
	initialized  bool
)

func GetServerConfig() *ServerConfig {
	if !initialized {
		panic("should initialize first!")
	}
	return serverConfig
}

func Initialize(isDebug bool) {

	// 获取当前工作目录
	rootDir, err := os.Getwd()
	if err != nil {
		logger.Panic("gwt wd error", err)
	}
	configFileName := "config-dev.yaml"
	if !isDebug {
		configFileName = "config-prod.yaml"
	}

	for i := 0; i < 5; i++ {
		if _, err := os.Stat(rootDir + "/" + "config-dev.yaml"); err == nil {
			break
		}
		rootDir = filepath.Dir(rootDir)
		logger.Debug("root dir", map[string]any{"dir": rootDir})
	}

	logger.Debug("root dir", map[string]any{"dir": rootDir})

	// 初始化配置文件
	v := viper.New()
	v.SetConfigFile(rootDir + "/" + configFileName)
	if err := v.ReadInConfig(); err != nil {
		logger.Panic("读取配置失败", err)
	}
	serverConfig = new(ServerConfig)
	if err := v.Unmarshal(serverConfig); err != nil {
		logger.Panic("解析配置失败", err)
	}
	logger.Info("server config:", map[string]any{"config": serverConfig})

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		logger.Info("配置文件修改了...", nil)
		if err := v.Unmarshal(serverConfig); err != nil {
			panic(err)
		}
		logger.Info("server config:", map[string]any{"config": serverConfig})
	})
	initialized = true
}
