package connector

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
)
type BaseConfig struct {
	Provider string`json:"provider"`
	Broker string`json:"broker"`
	Topic string`json:"topic"`
	Zookeeper string`json:"zookeeper"`
	Credential string`json:"credential"`
}

var colors = flag.Bool("colors", false, "enable or disable colors")
var broker = flag.String("brokerAddr", "localhost:9092", "Kafka Broker Address")
var topic = flag.String("TopicName", "localhost:9092", "Kafka Topic")
var provider = flag.String("CloudProvider", "google", "E.g:AWS,Azure,GoogleCloud")
var log aurora.Aurora
func InitFunc()  {
	flag.Parse()
	//Load Config
	LoadConfig()
	// Start connector
}
// 3 Way to load Config
func LoadConfig()  {
	if flag.Parsed() == false {
		return
	}
	_, err := FromFile()
	if err != nil {
		return
	}
	var bs *BaseConfig
	if bs.Zookeeper == "" {
		panic("Not Config")
	}
}
func (bs *BaseConfig)FromEnv()  {
	bs.Provider = os.Getenv("PROVIDER")
	bs.Broker = os.Getenv("BROKERS")
	bs.Zookeeper = os.Getenv("ZOOKEEPER")
	bs.Credential = os.Getenv("CRED")
	return
}
func FromFile() (bs *BaseConfig, err error) {
	var vbs *BaseConfig
	fopen, err := os.Open("./conf.json")
	if err != nil {
		log.BrightRed(fmt.Sprintf("Cannot readfile [ERROR]:%s", err.Error()))
		return nil, err
	}
	if err := json.NewDecoder(fopen).Decode(&vbs);err != nil {
		log.BrightRed(fmt.Sprintf("Cannot parsed [ERROR]:%s", err.Error()))
		return nil, err
	}
	bs.Credential = vbs.Credential
	bs.Zookeeper = vbs.Zookeeper
	bs.Broker = vbs.Broker
	bs.Provider = vbs.Provider
	return bs, nil
}