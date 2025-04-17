/*
 * I was initially planning on adding program version checks, then I decided no to inorder to make the command as fast
 * as possible. Which might have been a stupid idea
 */

package main

import (
	"fmt"
	"golang.design/x/clipboard"
	"log"
	"os"
	"strings"
)

func main() {
	error_ := clipboard.Init()

	if error_ != nil {
		log.Fatal(error_)
	}

	textArr := os.Args[1:]
	text := ""

	for i := 0; i < len(textArr); i++ {
		text += string(textArr[i]) + " "
	}

	if len(text) == 0 {
		log.Fatal("No text after the command!")
	}

	text = text[:len(text)-1]

	convertedText := ""
	caps := true

	for i := 0; i < len(text); i++ {
		text_ := text[i]

		if text_ >= 'A' && text_ <= 'Z' || text_ >= 'a' && text_ <= 'z' {
			if caps {
				convertedText += strings.ToUpper(string(text_))
			} else {
				convertedText += strings.ToLower(string(text_))
			}

			caps = !caps
		} else {
			convertedText += string(text_)
		}
	}

	clipboard.Write(clipboard.FmtText, []byte(convertedText))
	fmt.Println(convertedText)
	fmt.Println("Alternated the case and copied to clipboard!")
}
