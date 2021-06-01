package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

type Editor struct {
	CursorRow int
	CursorCol int

	Content []EditorChar
}

type EditorChar struct {
	Char    rune
	Index   int
	Color   string
	IsSpace bool
}

var editor *Editor = &Editor{}

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err.Error())
	}
	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			color.Red(err.Error())
		}

		if key == keyboard.KeyBackspace {
			editor.Content = editor.Content[:len(editor.Content)-1]
		} else if key == keyboard.KeySpace {
			char = ' '
		}

		if key != keyboard.KeyBackspace {
			var newChar EditorChar = EditorChar{
				Char: char,
			}

			editor.Content = append(editor.Content, newChar)
		}

		//os.Stdout.WriteString("\x1b[H") // reposition the cursor
		// os.Stdout.WriteString("\x1b[2J") // clear the screen
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		for i := 0; i < len(editor.Content); i++ {
			fmt.Print(string(editor.Content[i].Char))
		}
	}
}
