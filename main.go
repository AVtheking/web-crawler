package main

import (
	"fmt"
	"os"
)

func main() {
	//get the args from the command line
	args := os.Args
	argsWithoutProgramName := args[1:]
	if len(argsWithoutProgramName) < 1 {
		fmt.Println("No website provided ")
		os.Exit(1)
	}
	if len(argsWithoutProgramName) > 1 {
		fmt.Println("Too many arguments provided")
		os.Exit(1)
	}
	baseUrl := argsWithoutProgramName[0]
	fmt.Println("starting crawl of ", baseUrl	)

}
