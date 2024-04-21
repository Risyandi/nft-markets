package config

import (
	"github.com/spf13/viper"
)

type KeyViperConfig interface {
	GetBool(key string) bool
	GetInt(key string) int
	GetString(key string) string
	GetStringSlice(key string) []string
	GetUInt64(key string) uint64
	GetStringMap(key string) map[string]interface{}
	InitConfig()
}

type viperConfig struct{}

func (vr *viperConfig) InitConfig() {
	viper.SetConfigType(`json`)
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func (vr *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (vr *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (vr *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (vr *viperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func (vr *viperConfig) GetUInt64(key string) uint64 {
	return viper.GetUint64(key)
}

func (vr *viperConfig) GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func ConfigViper() KeyViperConfig {
	vr := &viperConfig{}
	vr.InitConfig()
	return vr
}
