package main

import (
	"flag"
	"log"
)

func main() {

	log.Println("Intiated main")
	/** Basic flag declarations are available for string, integer, and
	boolean options
	*/

	/** &sample string and &99 and &flase are default values for the below
	 */
	strFlag := flag.String("str1", "sample string", "Please give the string value")

	intFlag := flag.Int("integer", 99, "Please enter integer value")

	boolFlag := flag.Bool("bol", false, "please enter boolean value")

	var strvalue string
	flag.StringVar(&strvalue, "strv", "string var", "Please enter string var")

	/** Once all flags are declared, call flag.Parse() to execute
	the command-line parsing.
	*/
	flag.Parse()

	log.Printf("String : %s", *strFlag)
	log.Printf("Value: %d", *intFlag)
	log.Println("Bool", *boolFlag)
	log.Println("string", strvalue)

	/** To run the program
	go run Flags.go -strv=Hello -str1=Naresh -integer=1011 -bol=true
	*/

	log.Println("Tail args:", flag.Args())
	/**
	go run Flags.go -strv=Hello -str1=Naresh -integer=1011 -bol=true a1 a2 a3 a4
	*/
	/**
	output
	 Intiated main
	 String : Naresh
	 Value: 1011
	 Bool true
	 string Hello
	 Tail args: [a1 a2 a3 a4]
	*/
	log.Println("Tail args:", flag.Args()[1])
	// output : a2

}
