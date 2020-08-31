package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	MYSQLCONF MysqlConf `yaml:"mysql"`
}
type MysqlConf struct {
	Host      string `yaml:"host,omitempty"`
	User      string `yaml:"user,omitempty"`
	Pwd       string `yaml:"pwd,omitempty"`
	DB        string `yaml:"db,omitempty"`
	Port      string `yaml:"port,omitempty"`
	AuthTable string `yaml:"auth_table,omitempty"`
	AuthName  string `yaml:"auth_name,omitempty"`
	AuthPwd   string `yaml:"auth_pwd,omitempty"`
}

var conf *Conf = nil

func GetConf() *Conf {
	if conf == nil {
		content, err := ioutil.ReadFile("conf.yaml")
		if err != nil {
			log.Fatal(err)
		}
		cfg := &Conf{}
		// If the entire config body is empty the UnmarshalYAML method is
		// never called. We thus have to set the DefaultConfig at the entry
		// point as well.
		err2 := yaml.Unmarshal([]byte(string(content)), cfg)
		if err2 != nil {
			log.Fatal(err)
		}
		conf = cfg
	}
	return conf
}

func GetMysqlDataSourc() string {
	return GetConf().MYSQLCONF.User + ":" + GetConf().MYSQLCONF.Pwd + "@tcp(" + GetConf().MYSQLCONF.Host + ":" + GetConf().MYSQLCONF.Port + ")/" + GetConf().MYSQLCONF.DB
}

func GetAuthTableName() string {
	return GetConf().MYSQLCONF.AuthTable
}
func GetAuthName() string {
	return GetConf().MYSQLCONF.AuthName
}
func GetAuthPwd() string {
	return GetConf().MYSQLCONF.AuthPwd
}
