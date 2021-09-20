package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("welcome to stephano coin\r\n")
	fmt.Printf("Please use the following commands:\r\n")
	fmt.Printf("explorer: Start the HTML Explorer\r\n")
	fmt.Printf("rest: Start the REST API (recommended)\r\n")
	os.Exit(1)
}

func main() {

	if len(os.Args) < 2 {
		usage()
	}

	//flag 가 많다면, flagset 을 사용하는게 유용함
	//플래그에타입을 정할 수 있음
	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Printf("Start Explorer\r\n")
	case "rest":
		rest.Parse(os.Args[2:])
		//fmt.Println("Start REST API")
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(*portFlag)
		fmt.Println("Start server")
	}

	fmt.Println(*portFlag)

}
