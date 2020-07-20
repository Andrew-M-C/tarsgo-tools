package config

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/conf"
)

var once sync.Once
var globalError error
var tarsConfig *conf.Conf

type Config struct{}

func initialize() {
	// ensure that input parameter is parsed
	tars.GetServerConfig()

	// read config from tars
	tars_conf, err := conf.NewConf(tars.ServerConfigPath)
	if err != nil {
		globalError = fmt.Errorf("conf.NewConf error: %w", err)
		return
	}

	// done
	tarsConfig = tars_conf
	return
}

func NewConfig() (*Config, error) {
	once.Do(initialize)
	if nil == tarsConfig {
		return nil, globalError
	}
	return &Config{}, nil
}

func (c *Config) GetString(domain string, key string, dft ...string) (value string, exist bool) {
	theMap := tarsConfig.GetMap(domain)
	value, exist = theMap[key]
	if false == exist {
		if len(dft) > 0 {
			value = dft[0]
		}
	}
	return
}

func (c *Config) GetInt(domain string, key string, dft ...int) (ret int, exist bool) {
	var str string
	str, exist = c.GetString(domain, key)
	if exist {
		num, err := strconv.Atoi(str)
		if err != nil {
			exist = false
		} else {
			ret = num
		}
	}
	if false == exist {
		if len(dft) > 0 {
			ret = dft[0]
		} else {
			ret = 0
		}
	}
	return
}

func (c *Config) GetLong(domain string, key string, dft ...int64) (ret int64, exist bool) {
	var str string
	str, exist = c.GetString(domain, key)
	if exist {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			exist = false
		} else {
			ret = num
		}
	}
	if false == exist {
		if len(dft) > 0 {
			ret = dft[0]
		} else {
			ret = 0
		}
	}
	return
}

func (c *Config) GetUlong(domain string, key string, dft ...uint64) (ret uint64, exist bool) {
	var str string
	str, exist = c.GetString(domain, key)
	if exist {
		num, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			exist = false
		} else {
			ret = num
		}
	}
	if false == exist {
		if len(dft) > 0 {
			ret = dft[0]
		} else {
			ret = 0
		}
	}
	return
}
