package common

import (
	"fmt"
	"github.com/kataras/golog"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var G_AppConfig *AppConfig

type (

	AppConfig struct {
		DBPostgres DBPostgres  `yaml:"postgresql"`
	}

	DBPostgres struct {
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
		BdName string `yaml:"dbname"`
		Host string `yaml:"host"`
		Port int `yaml:"port"`
	}

)

/**
db url
 */
func (conf DBPostgres) DBUrl() string {
	/*return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", conf.UserName,
		conf.Password, conf.Host, conf.Port, conf.BdName, "utf8")*/
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", conf.Host, conf.Port, conf.UserName,
		 conf.BdName, conf.Password)
}

/**
读取yaml配置
 */
func ReadYamlConfig(path string) (*AppConfig, error)  {
	conf := &AppConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}

/**
初始化配置
 */
func InitConfig() {
	prefixPath, err := os.Getwd()
	if err != nil {
		golog.Fatalf("读取文件路径错误", err.Error())
	}
	tmpPrefixPath := strings.Replace(prefixPath, "\\", "/", -1)
	path := "/config/app.yaml"
	conf, err := ReadYamlConfig(tmpPrefixPath + path)
	if err != nil {
		golog.Fatalf("解析配置文件错误", err.Error())
	}
	G_AppConfig = conf
}