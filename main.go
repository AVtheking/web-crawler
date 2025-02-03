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

	const maxConcurrency = 10
	cfg, err := configure(baseUrl, maxConcurrency)
	if err != nil {
		fmt.Println("error configuring: ", err)
		os.Exit(1)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	for url, count := range cfg.pages {
		fmt.Printf("URL: %s, Count: %d\n", url, count)
	}

}
