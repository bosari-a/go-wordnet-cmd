package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	wordnet "github.com/bosari-a/go-wordnet"
)

// ansi color codes
const RED = "\x1b[31m"
const GREEN = "\x1b[32m"
const BLUE = "\x1b[34m"
const YELLOW = "\x1b[33m"
const PURPLE = "\x1b[35m"
const RESET = "\x1b[0m"

var pos = []string{"noun", "verb", "adj", "adv"}

func main() {
	if len(os.Args) != 2 {
		usageMsg := fmt.Sprintf("%vIncorrect usage%v:\n\tUsage: %vwordnet%v %vsome_word%v\n", RED, RESET, BLUE, RESET, PURPLE, RESET)
		log.Fatal(usageMsg)
	}
	dictpath := os.Getenv("WORDNET_DICT_PATH")
	var word wordnet.Word
	word.Word = os.Args[1]
	err := word.GetDefinitions(dictpath)
	if err != nil {
		log.Fatal(err)
	}
	output := fmt.Sprintf("%v%v%v:\n", PURPLE, word.Word, RESET)
	goos := runtime.GOOS
	if goos != "linux" {
		output = word.Word + ":\n"
	}
	for _, v := range word.Definitions {
		substr := strings.Split(v, " - ")
		label := substr[0]
		def := substr[1]
		if goos == "linux" {
			output += fmt.Sprintf("%v%v%v: %v\n\n", getColor(label), label, RESET, def)
		} else {
			output += fmt.Sprintf("%v: %v\n\n", label, def)
		}

	}
	fmt.Print(output)
}

func getColor(label string) string {
	trueLabel := strings.Split(label, "(")[0]
	switch trueLabel {
	case pos[0]:
		return GREEN
	case pos[1]:
		return RED
	case pos[2]:
		return BLUE
	case pos[3]:
		return YELLOW
	}
	return RESET
}
