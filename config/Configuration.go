package config

import (
	"flag"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Mysql struct {
		DriverName string `yaml:"driver_name"`
		Host       string `yaml:"mysql_host"`
		Username   string `yaml:"mysql_username"`
		Pswd       string `yaml:"mysql_pswd"`
		DbName     string `yaml:"mysql_dbName"`
	} `yaml:"mysql"`

	Kafka struct {
		Broker      string `yaml:"broker"`
		UserName    string `yaml:"user_name"`
		Pswd        string `yaml:"password"`
		GroupId     string `yaml:"group_id"`
		TopicPrefix string `yaml:"topic_prefix"`
	} `yaml:"kafka"`
}

var configuration *Configuration

func LoadConfiguration() error {
	configFilePath := flag.String("C", "config/conf.yaml", "config file path")
	flag.Parse()
	log.Println("@@@Loaded the configFilePath:", *configFilePath)
	data, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	if config.Mysql.Pswd == "" {
		config.Mysql.Pswd = os.Getenv("mysqlPSWD")
	}
	if config.Kafka.Pswd == "" {
		config.Kafka.Pswd = os.Getenv("kafkaPSWD")
	}
	configuration = &config
	return err
}

func GetConfig() *Configuration {
	return configuration
}
