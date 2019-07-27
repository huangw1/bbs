package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *Config

type Config struct {
	Env            string `yaml:"Env"`
	BaseUrl        string `yaml:"BaseUrl"`
	Port           int    `yaml:"Port"`
	ShowSql        bool   `yaml:"ShowSql"`
	ViewsPath      string `yaml:"ViewsPath"`
	RootStaticPath string `yaml:"RootStaticPath"`
	StaticPath     string `yaml:"StaticPath"`

	MySqlUrl  string `yaml:"MySqlUrl"`
	RedisAddr string `yaml:"RedisAddr"`

	Redis struct {
		Addr     string `yaml:"Addr"`
		Password string `yaml:"Password"`
	} `yaml:"Redis"`

	OauthServer struct {
		AuthUrl  string `yaml:"AuthUrl"`
		TokenUrl string `yaml:"TokenUrl"`
	} `yaml:"OauthServer"`

	OauthClient struct {
		ClientId          string `yaml:"ClientId"`
		ClientSecret      string `yaml:"ClientSecret"`
		ClientRedirectUrl string `yaml:"ClientRedirectUrl"`
		ClientSuccessUrl  string `yaml:"ClientSuccessUrl"`
	} `yaml:"OauthClient"`

	Github struct {
		ClientID     string `yaml:"ClientID"`
		ClientSecret string `yaml:"ClientSecret"`
	} `yaml:"Github"`

	AliyunOss struct {
		Host         string `yaml:"Host"`
		Bucket       string `yaml:"Bucket"`
		Endpoint     string `yaml:"Endpoint"`
		AccessId     string `yaml:"AccessId"`
		AccessSecret string `yaml:"AccessSecret"`
	} `yaml:"AliyunOss"`

	Smtp struct {
		Addr     string `yaml:"Addr"`
		Port     string `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
	} `yaml:"Smtp"`
}

func InitConfig(filename string) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Fatalf("init config failed: %s", err.Error())
		return
	}

	Conf = &Config{}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		logrus.Fatalf("unmarshal config failed: %s", err.Error())
		return
	}
}
