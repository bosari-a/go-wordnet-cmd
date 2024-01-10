package main

import (
	"fmt"
	"os"

	wordnet "github.com/bosari-a/go-wordnet"
)

const RED = "\x1b[31m"
const GREEN = "\x1b[32m"
const BLUE = "\x1b[34m"
const YELLOW = "\x1b[33m"
const RESET = "\x1b[0m"

func main() {
	dictpath := os.Getenv("WORDNET_DICT_PATH")
	var word wordnet.Word
	word.Word = os.Args[1]
	err := word.GetDefinitions(dictpath)
	if err != nil {
		panic(err)
	}
	output := fmt.Sprintf("%v%v%v:\n", GREEN, word.Word, RESET)
	for k, v := range word.Definitions {
		output += fmt.Sprintf("%v%v%v: %v\n", BLUE, k, RESET, v)
	}
	fmt.Print(output)
}
