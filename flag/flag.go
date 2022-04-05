package main

import (
	"flag"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var dbdetails = make(map[string]map[string]interface{})

func main() {

	log.Println("Intiated main")
	/** Basic flag declarations are available for string, integer, and
	boolean options
	*/

	/** Pass the environment "prod" or "dev" to commandline and defalt is "dev"
	 */
	env := flag.String("env", "dev", "Select environment")

	/** Pass the port
	 */
	// port := flag.Int("port", 8080, "Pass the port here")

	// dbflag = flag.Bool("dbflag", false, "Pass the db availability")

	// var temp string
	// flag.StringVar(&temp, "strv", "string var", "P")

	/** Once all flags are declared, call flag.Parse() to execute
	the command-line parsing.
	*/
	flag.Parse()

	log.Printf("String : %s", *env)
	readYaml()

	if *env == "dev" {
		db := dbdetails["dev"]
		log.Println("String : Host:", db["host"])
		log.Println("String : Port:", db["port"])
	} else {
		db := dbdetails["prod"]
		log.Println("String : Host:", db["host"])
		log.Println("String : Port:", db["port"])
	}

	/** To run the program
	commandline arguements with name
	*/
	log.Println("Tail args:", flag.Args())

	// receive the 2nd value from args.
	log.Println("Tail args:", flag.Args()[1])
	// output : a2

}
func readYaml() {
	yamlFile, err := ioutil.ReadFile("./flag/conf.yaml")

	if err != nil {
		log.Panic(err)
	}

	yaml.Unmarshal(yamlFile, dbdetails)

}
