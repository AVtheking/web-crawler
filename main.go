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
		fmt.Println("no website provided ")
		os.Exit(1)
	}
	if len(argsWithoutProgramName) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := argsWithoutProgramName[0]
	fmt.Println("starting crawl of ", baseUrl)

	pages := make(map[string]int)

	pages = crawlPage(baseUrl, baseUrl, pages)
	for url, count := range pages {
		fmt.Printf("URL: %s, Count: %d\n", url, count)
	}

}
