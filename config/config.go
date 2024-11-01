package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

type MainConfig struct {
	Server      ServerConfig `yaml:"Server"`
	Database    DBConfig     `yaml:"Database"`
	InternalAPI APIConfig    `yaml:"InternalAPI"`
}

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

var env string

type (
	APIConfig struct {
		BasePath      string `yaml:"BasePath"`
		APITimeout    int    `yaml:"APITimeout"`
		EnableSwagger bool   `yaml:"EnableSwagger" default:"true"`
	}

	ServerConfig struct {
		Port            string        `yaml:"Port"`
		ConsumerPort    string        `yaml:"ConsumerPort"`
		GracefulTimeout time.Duration `yaml:"GracefulTimeout"`
		ReadTimeout     time.Duration `yaml:"ReadTimeout"`
		WriteTimeout    time.Duration `yaml:"WriteTimeout"`
	}

	DBConfig struct {
		SlaveDSN        string `yaml:"SlaveDSN"`
		MasterDSN       string `yaml:"MasterDSN"`
		RetryInterval   int    `yaml:"RetryInterval"`
		MaxIdleConn     int    `yaml:"MaxIdleConn"`
		MaxConn         int    `yaml:"MaxConn"`
		ConnMaxLifetime string `yaml:"ConnMaxLifetime"`
	}
)

func ReadModuleConfig(cfg interface{}, module, configLocation string) interface{} {
	if configLocation == "" {
		configLocation = "config/files"
	}

	err := ReadModuleConfigFile(cfg, configLocation, module)
	if err != nil {
		log.Fatalf("failed to read config for %v err:%v", module, err.Error())
	}
	return cfg
}

func init() {
	env = os.Getenv("ENV")
	if env == "" {
		env = EnvDevelopment
	}
}

func ReadModuleConfigFile(cfg interface{}, path string, module string) error {
	environ := env
	getFormatFile := filePath(path)

	switch getFormatFile {
	case ".json":
		fname := path + "/" + module + "." + environ + ".json"
		jsonFile, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonFile, cfg)
	default:
		fname := path + "/" + module + "." + environ + ".yaml"
		yamlFile, err := ioutil.ReadFile(fname)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(yamlFile, cfg)
	}

}

func filePath(root string) string {
	var file string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		file = filepath.Ext(info.Name())
		return nil
	})
	return file
}
