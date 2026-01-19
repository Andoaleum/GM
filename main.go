package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if(len(os.Args) < 3){
		fmt.Println("Usage: [filename] [show score](optional-false)")
		log.Fatal("Missing arguments!")

		return
	}
	source, err := os.ReadFile(os.Args[1])
	if err != nil{log.Fatal(err)}
	showScore := false;
	if(os.Args[2] == "true"){showScore = true}

	fmt.Println(censorBadWord(string(source),showScore))

}
