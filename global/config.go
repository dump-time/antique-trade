package global

import (
	"os"

	"github.com/dump-time/antique-trade/log"
	"gopkg.in/yaml.v2"
)

type configType struct {
	Serv struct {
		Addr           string
		TrustedProxies []string `yaml:"proxies"`
	}
	DB struct {
		Hostname string `yaml:"host"`
		Port     string
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		DBName   string `yaml:"name"`
	}
	Redis struct {
		Hostname string `yaml:"host"`
		Password string `yaml:"pass"`
		Port     string `yaml:"port"`
		Secret   string `yaml:"secret"`
	}
}

var Config configType

func initConfig() {
	configFile, err := os.ReadFile(CmdOpts.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatal(err)
	}
	log.Info(Config)
}
