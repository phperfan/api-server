package config

import (
	"api-server/pkg/log"
	"strings"

	"github.com/fsnotify/fsnotify"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config 读取配置
type Config struct {
	// Name 配置文件路径
	Name string
}

func (cfg Config) initConfig() error {
	if cfg.Name != "" {
		viper.SetConfigFile(cfg.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config.local")
	}
	viper.SetConfigType("yaml") // 设置配置文件格式为YAML
	viper.AutomaticEnv()        // 读取匹配的环境变量
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return errors.WithStack(err)
	}

	return nil
}

func (cfg Config) initLog() {
	if err := log.InitLogger(); err != nil {
		panic(err)
	}
	log.Al.Debug("初始化日志成功")
}

func (cfg Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Al.Infof("Config file changed: %s", e.Name)
	})
}

//Init 初始化配置文件
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	// 监控配置文件变化并热加载程序
	c.watchConfig()
	return nil
}
