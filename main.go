package main

import (
	"fmt"
	"os"
	"strings"
	
	//"github.com/alasher/d2reel/parser"
	"github.com/alasher/d2reel/fetcher"
)

const VERSION = "1.0.0"

func main() {
	
	if len(os.Args) < 2 {
		fmt.Println("No armguments, brother! Run 'd2reel help' for argument options.")
		os.Exit(0)
	}
	
	var ids []string
	
	for _, v := range os.Args[1:] {
		switch strings.ToLower(v) {
		case "help":
			printHelpOptions()
			os.Exit(0)
		case "--version":
			printVersion()
			os.Exit(0)
		default:
			ids = append(ids, v)
		}
	}
	
	fetcher.Fetch(ids)
	//parser.Parse(ids)
	
}

func printHelpOptions() {
	fmt.Print("Hey, friend.\nd2reel match_id+ [flag option]*\n\n")
	fmt.Print("You can place any number of match ids, d2reel will parse them all.\nFlags can appear anywhere in the argument list, not just after all the match ids.\n\n");
	
	// Aren't any real options right now... I'll add more once I get the ball rolling with development.
	fmt.Print("Flag options:\n")
	fmt.Printf("%10s - Brings up this help menu. You must have used it at least once now. Nice!\n", "help")
	fmt.Printf("%10s - Displays the current d2reel version.\n", "--version")
}

func printVersion() {
	fmt.Printf("You're running d2reel v%v\n", VERSION)
}