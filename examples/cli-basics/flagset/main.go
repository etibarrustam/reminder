package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no comment")
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI BASICS", "greet msg")
		err := greetCmd.Parse(os.Args[2:])

		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("hello %s\n", *msgFlag)
	case "help":
		fmt.Println("help msg")
	default:
		fmt.Printf("unknown command %s\n", cmd)
	}
}
