package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var data = make(map[interface{}]interface{})

func main() {

	log.Println("App Initiated")
	log.Println("Data:: ", readYaml())
	log.Println("Data:: Name:: %s, \n Hits:: %d, \n Time:: %d", readYaml()["name"], readYaml()["hits"], readYaml()["time"])
}
func readYaml() map[interface{}]interface{} {
	yamlFile, err := ioutil.ReadFile("./yaml/conf.yaml")

	if err != nil {
		log.Panic(err)
	}

	yaml.Unmarshal(yamlFile, data)
	return data

}
