/*
 * I was initially planning on adding program version checks, then I decided no to inorder to make the command as fast
 * as possible. Which might have been a stupid idea
 */

package main

import (
	"os"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

func main() {
	clipboardWorks := true

	err := clipboard.Init()

	if err != nil {
		clipboardWorks = false
		println("Clipboard init failed.")
	}

	textArr := os.Args[1:]
	text := ""

	for i := 0; i < len(textArr); i++ {
		text += string(textArr[i]) + " "
	}

	if len(text) == 0 {
		panic("No text after the command!")
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

	println(convertedText)

	if clipboardWorks {
		println("Waiting for a second so it copies to your clipboard...")
		clipboard.Write(clipboard.FmtText, []byte(convertedText))
		time.Sleep(1 * time.Second)
		println("Copied to clipboard!")
	}
}
