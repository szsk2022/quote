package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var C Conf

func init() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("read file err %v\n", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &C)
	if err != nil {
		log.Fatalf("yaml unmarshal: %v\n", err)
		return
	}
}
