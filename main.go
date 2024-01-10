package main

import (
	"fmt"
	"os"
	"strings"

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
	for _, v := range word.Definitions {
		substr := strings.Split(v, " - ")
		output += fmt.Sprintf("%v%v%v: %v\n\n", BLUE, substr[0], RESET, substr[1])
	}
	fmt.Print(output)
}
