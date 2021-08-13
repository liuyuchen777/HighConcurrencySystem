/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 03:37:38
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 08:33:34
 * @Description:
 * @FilePath: /super_high_concurrency_system/config/conf.go
 * @GitHub: https://github.com/liuyuchen777
 */
package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const configFilePath string = "./config-dev.yaml"

var configFile []byte

type AppConfig struct {
	App App `yaml:"app"`
}

type App struct {
	Database        Database `yaml:"database"`
	Redis           Redis    `yaml:"redis"`
	FlushAllForTest bool     `yaml:"flushAllForTest"`
}

type Database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
	Address  string `yaml:"address"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

type Redis struct {
	Address     string `yaml:"address"`
	Network     string `yaml:"network"`
	Password    string `yaml:"password"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxActive   int    `yaml:"maxActive"`
	IdleTimeout int    `yaml:"idleTimeout"`
}

func GetAppConfig() (appConfig AppConfig, err error) {
	err = yaml.Unmarshal(configFile, &appConfig)
	return appConfig, err
}

func init() {
	/* 自动执行 */
	var err error
	configFile, err = ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("yamlFile. Get err %v", err)
	}
}
