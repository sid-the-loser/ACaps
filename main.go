package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"log"
	"os"
	"strings"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}

	textArr := os.Args[1:]
	text := ""

	for i := 0; i < len(textArr); i++ {
		text += string(textArr[i]) + " "
	}

	text = text[:len(text)-1]

	convertedText := ""
	caps := true

	for i := 0; i < len(text); i++ {
		text_ := string(text[i])
		if caps {
			convertedText += strings.ToUpper(text_)
		} else {
			convertedText += strings.ToLower(text_)
		}
		caps = !caps
	}

	clipboard.Write(clipboard.FmtText, []byte(convertedText))
	fmt.Println(convertedText)
	fmt.Println("Alternated the case and copied to clipboard!")
}
