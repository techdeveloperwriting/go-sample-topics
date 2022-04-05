package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var data = make(map[string]map[string]interface{})

func main() {

	log.Println("App Initiated")
	data = readYaml()
	log.Println("Data:: ", data)

	//data := make(map[interface{}]interface{})
	data1 := readYaml()["mongo"]
	log.Println("Data:: ", data1["name"])
	//	log.Println("Data:: Name:: %s, \n Hits:: %d, \n Time:: %d", readYaml()["name"], readYaml()["hits"], readYaml()["time"])
}
func readYaml() map[string]map[string]interface{} {
	yamlFile, err := ioutil.ReadFile("./yaml/conf_copy.yaml")

	if err != nil {
		log.Panic(err)
	}

	yaml.Unmarshal(yamlFile, data)
	return data

}
