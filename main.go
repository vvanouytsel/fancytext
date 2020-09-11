package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput() string {
	input := os.Args[1:]
	if len(input) != 1 {
		err := help()
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Converting '" + input[0] + "' to fancy text.")
	return input[0]
}

func help() error {
	fmt.Println("Example: go run . \"Convert this using Fancy Text\"")
	return (errors.New("help: please specify one string that needs to be converted"))
}

func convertFancy(text string) string {
	var fancyText string
	fancyLetters := map[string]string{
		"a": "4",
		"b": "8",
		"c": "c",
		"d": "d",
		"e": "3",
		"f": "f",
		"g": "g",
		"h": "h",
		"i": "1",
		"j": "j",
		"k": "k",
		"l": "l",
		"m": "m",
		"n": "n",
		"o": "0",
		"p": "p",
		"q": "q",
		"r": "r",
		"s": "s",
		"t": "t",
		"u": "u",
		"v": "v",
		"w": "w",
		"x": "x",
		"y": "y",
		"z": "z",
		" ": " ",
	}

	for _, char := range text {
		// Range returns a rune (which is the ASCII representation of the char)
		// We have to convert it to a string if we want the char based value
		// fmt.Println(string(char), "--->", fancyLetters[string(char)])
		_, ok := fancyLetters[string(char)]
		if !ok {
			// Could not find key in map
			err := errors.New("convertFancy: could not find key '" + string(char) + "' in map")
			log.Fatal(err)
		}
		fancyText += fancyLetters[string(char)]
	}

	return fancyText
}

func main() {
	input := getInput()
	fmt.Println(convertFancy(strings.ToLower(input)))
}
