package config

import (
	"strings"
)

const (
	mode_production  = "PRODUCTION"
	mode_test        = "TEST"
	mode_development = "DEVELOPMENT"

	LOG_LEVEL_DEBUG = "DEBUG"
	LOG_LEVEL_INFO  = "INFO"
	LOG_LEVEL_WARN  = "WARN"
	LOG_LEVEL_ERROR = "ERROR"
)

type config struct {
	Mode     string      `yaml:"mode"`
	LogLevel string      `yaml:"logLevel"`
	HTTP     *httpConfig `yaml:"http"`
	GRPC     *grpcConfig `yaml:"grpc"`
}

type httpConfig struct {
	Enable bool `yaml:"enable"`
	Debug  bool `yaml:"debug"`
	Port   int  `yaml:"port"`
}

type grpcConfig struct {
	Enable bool `yaml:"enable"`
	Debug  bool `yaml:"debug"`
	Port   int  `yaml:"port"`
}

func (c *config) init() error {
	c.Mode = strings.ToUpper(c.Mode)
	if c.Mode != mode_production && c.Mode != mode_development && c.Mode != mode_test {
		c.Mode = mode_production
	}

	c.LogLevel = strings.ToUpper(c.LogLevel)
	if c.LogLevel != LOG_LEVEL_DEBUG && c.LogLevel != LOG_LEVEL_INFO && c.LogLevel != LOG_LEVEL_WARN && c.LogLevel != LOG_LEVEL_ERROR {
		c.LogLevel = LOG_LEVEL_INFO
	}

	if c.HTTP == nil {
		c.HTTP = &httpConfig{Enable: true}
	}
	if c.GRPC == nil {
		c.GRPC = &grpcConfig{Enable: true}
	}

	if err := c.HTTP.init(); err != nil {
		return err
	}

	if err := c.GRPC.init(); err != nil {
		return err
	}

	return nil
}

func (c *httpConfig) init() error {
	if !c.Enable {
		return nil
	}
	if c.Port <= 0 {
		c.Port = 10021
	}
	return nil
}

func (c *grpcConfig) init() error {
	if !c.Enable {
		return nil
	}
	if c.Port <= 0 {
		c.Port = 10022
	}
	return nil
}

func (c config) IsProduction() bool {
	return c.Mode == mode_production
}

func (c config) IsTest() bool {
	return c.Mode == mode_test
}

func (c config) IsDevelopment() bool {
	return c.Mode == mode_development
}
