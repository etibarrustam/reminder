package main

import (
	"flag"
	"fmt"
	"os"
	"reminder/client"
)

var (
	backendURIFlag = flag.String("backend", "http://localhost:8080", "Backend api url")
	helpFlag       = flag.Bool("help", false, "Display helpful message")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendURIFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()

	if err != nil {
		fmt.Printf("switch error: %v\n", err)

		os.Exit(2)
	}
}
