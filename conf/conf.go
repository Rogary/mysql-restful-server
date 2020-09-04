package conf

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	MYSQLCONF mysqlConf `yaml:"mysql"`
}
type mysqlConf struct {
	Host       string `yaml:"host,omitempty"`
	User       string `yaml:"user,omitempty"`
	Pwd        string `yaml:"pwd,omitempty"`
	DB         string `yaml:"db,omitempty"`
	Port       string `yaml:"port,omitempty"`
	EnableAuth string `yaml:"enable_auth,omitempty"`
	AuthTable  string `yaml:"auth_table,omitempty"`
	AuthName   string `yaml:"auth_name_field,omitempty"`
	AuthPwd    string `yaml:"auth_pwd_field,omitempty"`
}

var confs *conf

func getConf() *conf {
	if confs == nil {
		content, err := ioutil.ReadFile("conf.yaml")
		if err != nil {
			log.Fatal(err)
		}
		cfg := &conf{}
		// If the entire config body is empty the UnmarshalYAML method is
		// never called. We thus have to set the DefaultConfig at the entry
		// point as well.
		err2 := yaml.Unmarshal([]byte(string(content)), cfg)
		if err2 != nil {
			log.Fatal(err)
		}
		confs = cfg
	}
	return confs
}

// GetMysqlDataSource get datasourceUrl
func GetMysqlDataSource() string {
	return getMysqlUser() + ":" + getMysqlPwd() + "@tcp(" + getMysqlHost() + ":" + getMysqlPort() + ")/" + getMysqlDb()
}

func getMysqlUser() string {
	mysqlUser := os.Getenv("MYSQL_USER")
	if mysqlUser != "" {
		return mysqlUser
	}
	return getConf().MYSQLCONF.User
}
func getMysqlPwd() string {
	mysqlPwd := os.Getenv("MYSQL_PWD")
	if mysqlPwd != "" {
		return mysqlPwd
	}
	return getConf().MYSQLCONF.Pwd
}

func getMysqlHost() string {
	mysqlHost := os.Getenv("MYSQL_HOST")
	if mysqlHost != "" {
		return mysqlHost
	}
	return getConf().MYSQLCONF.Host
}

func getMysqlPort() string {
	mysqlPort := os.Getenv("MYSQL_PORT")
	if mysqlPort != "" {
		return mysqlPort
	}
	return getConf().MYSQLCONF.Port
}
func getMysqlDb() string {
	mysqlDb := os.Getenv("MYSQL_DB")
	if mysqlDb != "" {
		return mysqlDb
	}
	return getConf().MYSQLCONF.DB
}

// GetAuthTableName get auth table name
func GetAuthTableName() string {
	mysqlAuthTable := os.Getenv("MYSQL_AUTH_TABLE")
	if mysqlAuthTable != "" {
		return mysqlAuthTable
	}
	return getConf().MYSQLCONF.AuthTable
}

// GetAuthName get GetAuthNameField
func GetAuthName() string {
	mysqlAuthNameField := os.Getenv("MYSQL_AUTH_NAME_FIELD")
	if mysqlAuthNameField != "" {
		return mysqlAuthNameField
	}
	return getConf().MYSQLCONF.AuthName
}

// GetEnableAuth get GetEnableAuth
func GetEnableAuth() bool {
	mysqlEnableAuth := os.Getenv("MYSQL_ENABLE_AUTH")
	if mysqlEnableAuth != "" {
		return mysqlEnableAuth == "true"
	}
	return getConf().MYSQLCONF.EnableAuth == "true"
}

// GetAuthPwd get GetAuthPwdField
func GetAuthPwd() string {
	mysqlAuthPwdField := os.Getenv("MYSQL_AUTH_PWD_FIELD")
	if mysqlAuthPwdField != "" {
		return mysqlAuthPwdField
	}
	return getConf().MYSQLCONF.AuthPwd
}
