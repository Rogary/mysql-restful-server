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
type memoryConf struct {
	MysqlHost     string
	MysqlUser     string
	MysqlPwd      string
	MysqlDb       string
	MysqlPort     string
	EnableAuth    bool
	AuthTable     string
	AuthFieldName string
	AuthFieldPwd  string
}

var mConf = &memoryConf{"", "", "", "", "", false, "", "", ""}
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
	if mConf.MysqlUser == "" {
		mysqlUser := os.Getenv("MYSQL_USER")
		if mysqlUser != "" {
			mConf.MysqlUser = mysqlUser
		}
		mConf.MysqlUser = getConf().MYSQLCONF.User
	}
	return mConf.MysqlUser
}
func getMysqlPwd() string {
	if mConf.MysqlPwd == "" {
		mysqlPwd := os.Getenv("MYSQL_PWD")
		if mysqlPwd != "" {
			mConf.MysqlPwd = mysqlPwd
		}
		mConf.MysqlPwd = getConf().MYSQLCONF.Pwd
	}
	return mConf.MysqlPwd
}

func getMysqlHost() string {
	if mConf.MysqlHost == "" {
		mysqlHost := os.Getenv("MYSQL_HOST")
		if mysqlHost != "" {
			mConf.MysqlHost = mysqlHost
		}
		mConf.MysqlHost = getConf().MYSQLCONF.Host
	}
	return mConf.MysqlHost
}

func getMysqlPort() string {
	if mConf.MysqlPort == "" {
		mysqlPort := os.Getenv("MYSQL_PORT")
		if mysqlPort != "" {
			mConf.MysqlPort = mysqlPort
		}
		mConf.MysqlPort = getConf().MYSQLCONF.Port
	}
	return mConf.MysqlPort
}
func getMysqlDb() string {
	if mConf.MysqlDb == "" {
		mysqlDb := os.Getenv("MYSQL_DB")
		if mysqlDb != "" {
			mConf.MysqlDb = mysqlDb
		}
		mConf.MysqlDb = getConf().MYSQLCONF.DB
	}
	return mConf.MysqlDb
}

// GetAuthTableName get auth table name
func GetAuthTableName() string {
	if mConf.AuthTable == "" {

		mysqlAuthTable := os.Getenv("MYSQL_AUTH_TABLE")
		if mysqlAuthTable != "" {
			mConf.AuthTable = mysqlAuthTable
		}
		mConf.AuthTable = getConf().MYSQLCONF.AuthTable
	}
	return mConf.AuthTable
}

// GetAuthName get GetAuthNameField
func GetAuthName() string {
	if mConf.AuthFieldName == "" {
		mysqlAuthNameField := os.Getenv("MYSQL_AUTH_NAME_FIELD")
		if mysqlAuthNameField != "" {
			mConf.AuthFieldName = mysqlAuthNameField
		}
		mConf.AuthFieldName = getConf().MYSQLCONF.AuthName
	}
	return mConf.AuthFieldName
}

// GetEnableAuth get GetEnableAuth
func GetEnableAuth() bool {
	if mConf.EnableAuth == false {

		mysqlEnableAuth := os.Getenv("MYSQL_ENABLE_AUTH")
		if mysqlEnableAuth != "" {
			mConf.EnableAuth = mysqlEnableAuth == "true"
		}
		mConf.EnableAuth = getConf().MYSQLCONF.EnableAuth == "true"
	}
	return mConf.EnableAuth
}

// GetAuthPwd get GetAuthPwdField
func GetAuthPwd() string {
	if mConf.AuthFieldPwd == "" {

		mysqlAuthPwdField := os.Getenv("MYSQL_AUTH_PWD_FIELD")
		if mysqlAuthPwdField != "" {
			mConf.AuthFieldPwd = mysqlAuthPwdField
		}
		mConf.AuthFieldPwd = getConf().MYSQLCONF.AuthPwd
	}
	return mConf.AuthFieldPwd
}
