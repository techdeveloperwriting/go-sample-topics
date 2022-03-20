package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Define struct for reading yaml details
type Config struct {
	Hits int    `yaml:"hits"`
	Time int    `yaml:"time"`
	Name string `yaml:"name"`
}

func main() {
	log.Printf("App Initiated")
	// Using Struct
	var conf Config

	log.Println("Config struct::", conf.readYaml())
	log.Printf("Config details:: \n Name:: %s,\n Time:: %d,\n Hits:: %d",
		conf.readYaml().Name, conf.readYaml().Time, conf.readYaml().Hits)

}

func (conf *Config) readYaml() *Config {
	// Read yaml file from path
	yamlFile, err := ioutil.ReadFile("./yaml/conf.yaml")

	if err != nil {
		log.Panic(err)
	}
	// Convert byte array to struct
	err1 := yaml.Unmarshal(yamlFile, conf)

	if err1 != nil {
		log.Panic(err)
	}
	return conf

}
