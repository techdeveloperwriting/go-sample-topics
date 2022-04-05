package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Env struct {
	Env []EnvDetails `json:"environment"`
}

type EnvDetails struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	log.Println("App initiated")
	jsonFile, err := os.Open("json/config.json")
	if err != nil {
		panic(err)
	}
	temp := Env{}

	jsonByte, _ := ioutil.ReadAll(jsonFile)
	err1 := json.Unmarshal(jsonByte, &temp)
	if err1 != nil {
		panic(err)
	}
	log.Println("Hello ", temp)
	log.Println("Hello ", temp.Env[0])
	log.Println("Hello length", len(temp.Env))
	for i := 0; i < len(temp.Env); i++ {
		log.Println("Host ", temp.Env[i].Host)
		log.Println("Port ", temp.Env[i].Port)
	}
	defer jsonFile.Close()
}
