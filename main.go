package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//get the args from the command line
	args := os.Args
	argsWithoutProgramName := args[1:]
	if len(argsWithoutProgramName) < 1 {
		fmt.Println("no website provided ")
		os.Exit(1)
	}
	if len(argsWithoutProgramName) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := argsWithoutProgramName[0]
	maxConcurrency := argsWithoutProgramName[1]
	maxPages := argsWithoutProgramName[2]
	fmt.Println("starting crawl of ", baseUrl)

	maxConcurrencyInt, err := strconv.Atoi(maxConcurrency)
	if err != nil {
		fmt.Println("error converting maxConcurrency to int: ", err)
		os.Exit(1)
	}

	maxPagesInt, err := strconv.Atoi(maxPages)
	if err != nil {
		fmt.Println("error converting maxPages to int: ", err)
		os.Exit(1)
	}

	cfg, err := configure(baseUrl, maxConcurrencyInt, maxPagesInt)
	if err != nil {
		fmt.Println("error configuring: ", err)
		os.Exit(1)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	printReport(cfg.pages, baseUrl)

}
