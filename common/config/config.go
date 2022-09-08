package config

import (
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/spf13/viper"
)

var (
	Config config
	once   sync.Once
)

type config struct {
	Consul struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	WServer struct {
		Host  string `yaml:"host"`
		Port  int    `yaml:"port"`
		Check struct {
			Interval                       int `yaml:"interval"`
			DeregisterCriticalServiceAfter int `yaml:"deregisterCriticalServiceAfter"`
		}
	}

	AccountServer struct {
		Addr string `yaml:"addr"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	Github struct {
		Oauth struct {
			ClientId     string `yaml:"clientId"`
			ClientSecret string `yaml:"clientSecret"`
			RedirectUrl  string `yaml:"RedirectUrl"`
		}
	}

	Log struct {
		Level int `yaml:"level"`
	}
}

func getCurrentABPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Get curr path error")
	}
	abPath = path.Dir(filename)
	return abPath
}

func init() {
	once.Do(func() {
		confPath := os.Getenv("CONF_PATH")
		if confPath == "" {
			confPath = path.Join(getCurrentABPath(), "../../conf")
		}
		env := os.Getenv("ENV_MODE")
		if env == "" {
			env = "dev"
		}
		viper.AddConfigPath(confPath + "/" + env + "/")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		if err = viper.Unmarshal(&Config); err != nil {
			panic(err)
		}
	})

}
