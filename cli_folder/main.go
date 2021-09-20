package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("welcome to stephano coin\r\n")
	fmt.Printf("Please use the following flags:\r\n")
	fmt.Printf("-port=4000: Set the PORT of the server\r\n")
	fmt.Printf("-mode=rest: Choose between 'html' and 'rest'\r\n")
	os.Exit(1)
}

func main() {

	//no command just flag
	//default flag is in, so we don't need this exception
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest' ")

	//1번 argument 부터 쭉 확인
	flag.Parse()

	switch *mode {
	case "rest":
		fmt.Printf("Rest\r\n")
		//start rest api
		//rest.Start(*port)
	case "html":
		fmt.Printf("html\r\n")
		//start html explorer
		//explorer.Start(*port)

	default:
		usage()
	}

	fmt.Println(*port, *mode)

}
