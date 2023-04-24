package config

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/niubir/utils"
	yaml "gopkg.in/yaml.v2"
)

var (
	configDir      = "/config"
	configFilename = "config.yml"
	Config         config
)

func init() {
	if envConfigDir := os.Getenv(`CONFIG_DIR`); envConfigDir != "" {
		configDir = envConfigDir
	}
	if !utils.FilepathExist(configDir) {
		if err := os.Mkdir(configDir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	if envConfigFilename := os.Getenv(`CONFIG_FILENAME`); envConfigFilename != "" {
		configFilename = envConfigFilename
	}

	utils.BoldPrint("***CONFIG FILE***\n%s", configDir+"/"+configFilename)
}

// Init 初始化配置文件
func Init() error {
	data, err := ioutil.ReadFile(configDir + "/" + configFilename)
	if err != nil {
		return errors.New("invalid load config file: " + err.Error())
	}
	if err := yaml.Unmarshal(data, &Config); err != nil {
		return errors.New("unmarshal config file: " + err.Error())
	}
	if err := Config.init(); err != nil {
		return errors.New("failed to init config: " + err.Error())
	}

	configBody, _ := yaml.Marshal(Config)
	utils.BoldPrint("***CONFIG INFORMATION***\n\n%s", string(configBody))

	return nil
}
