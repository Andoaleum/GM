package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	sourcePath := "./source.txt"

	banPath := "./ban.txt"


	showScore := false

	for i := 1; i < len(os.Args); i++ {

		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println("Usage: $[source.txt] %[ban.txt] (optional)[show score--false]")
			return
		} else if strings.Contains(os.Args[i], "!") {
			sourcePath = strings.ReplaceAll(os.Args[i], "!", "")
		} else if strings.Contains(os.Args[i], "%") {
			banPath = strings.ReplaceAll(os.Args[i], "%", "")
		} else if os.Args[i] == "true" {
			showScore = true
		}
	}
	source, err := os.ReadFile(sourcePath)
	ban, err1 := os.ReadFile(banPath)
	bannedWords = strings.Fields(string(ban))
	if err != nil {
		log.Fatal(err)
	}
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(censorBadWord(string(source), showScore))
}
