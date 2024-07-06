package vip

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

const (
	ErrReadConfig = "error reading config: "
	StatusProd    = "production"
	StatusDev     = "development"
)

const (
	C_App      = "app" + "."
	C_Database = "database" + "."
)

type (
	VipConf struct{}
	VipRes  struct {
		AppName       string
		AppPort       int
		AppPortAsset  int
		AppHost       string
		AppToken      string
		AppStatus     string
		DbName        string
		DbPassword    string
		DbUsername    string
		DbPort        int
		DbHost        string
		DbMaxCon      int
		DbMaxLifetime int
		DbSchema      string
		PasetoSecret  string
		PasetExp      int
		ApiKey        string
		XTokenKey     string
		EmailPassword string
		EmailPort     string
		EmailHost     string
		EmailSender   string
	}
)

func New() *VipConf {
	return &VipConf{}
}

func (v *VipConf) config() (*viper.Viper, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, errors.New("error reading workdir: " + err.Error())
	}
	config := viper.New()
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(wd)
	errRead := config.ReadInConfig()
	if errRead != nil {
		return nil, errors.New("error reading yaml: " + errRead.Error())
	}
	return config, nil
}

func (v *VipConf) App() (*VipRes, error) {
	config, errConf := v.config()
	if errConf != nil {
		return nil, errConf
	}

	res := &VipRes{
		AppName:       config.GetString(C_App + "name"),
		AppPort:       config.GetInt(C_App + "port"),
		AppPortAsset:  config.GetInt(C_App + "port_asset"),
		AppHost:       config.GetString(C_App + "host"),
		AppToken:      config.GetString(C_App + "token"),
		AppStatus:     config.GetString(C_App + "status"),
		DbName:        config.GetString(C_Database + "name"),
		DbPassword:    config.GetString(C_Database + "password"),
		DbUsername:    config.GetString(C_Database + "username"),
		DbPort:        config.GetInt(C_Database + "port"),
		DbHost:        config.GetString(C_Database + "host"),
		DbMaxCon:      config.GetInt(C_Database + "max_con"),
		DbMaxLifetime: config.GetInt(C_Database + "max_lifetime"),
		DbSchema:      config.GetString(C_Database + "schema"),
		PasetoSecret:  config.GetString(C_App + "paseto_secret"),
		PasetExp:      config.GetInt(C_App + "paseto_exp"),
		ApiKey:        config.GetString(C_App + "api_key"),
		XTokenKey:     config.GetString(C_App + "x_token_key"),
		EmailPassword: config.GetString(C_App + "email_password"),
		EmailHost:     config.GetString(C_App + "email_host"),
		EmailPort:     config.GetString(C_App + "email_port"),
		EmailSender:   config.GetString(C_App + "email_sender"),
	}
	return res, nil
}
