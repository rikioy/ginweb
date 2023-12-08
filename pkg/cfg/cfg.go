package cfg

import (
	"github.com/fsnotify/fsnotify"
	"github.com/rikioy/ginweb/pkg/env"
	"github.com/spf13/viper"
)

const (
	GINWEB_FILENAME = "ginweb"
)

var (
	cfg *Config
)

type Config struct {
	viper *viper.Viper
	env   string
}

func init() {
	viper.SetConfigName(GINWEB_FILENAME)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	cfg = &Config{viper: viper.GetViper(), env: env.GetEnvString()}
}

func ReadInConfig() error {
	if cfg.env != "" {
		cfg.viper.SetConfigName(GINWEB_FILENAME + "_" + cfg.env)
	}
	return cfg.viper.ReadInConfig()
}

func OnConfigChange(onConfigChangeFunc func(e fsnotify.Event)) {
	cfg.viper.OnConfigChange(onConfigChangeFunc)
}

func WatchConfig() {
	cfg.viper.WatchConfig()
}

func SetEnv(env string) {
	cfg.env = env
}

func SetConfigName(name string) {
	cfg.viper.SetConfigName(name)
}

func AddConfigPath(path string) {
	cfg.viper.AddConfigPath(path)
}

func GetString(key string) string {
	return cfg.viper.GetString(key)
}

func GetInt(key string) int {
	return cfg.viper.GetInt(key)
}

func GetInt32(key string) int32 {
	return cfg.viper.GetInt32(key)
}

func GetInt64(key string) int64 {
	return cfg.viper.GetInt64(key)
}

func GetBool(key string) bool {
	return cfg.viper.GetBool(key)
}
