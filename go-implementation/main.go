package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

// Handle user input
func handleInput(root *node) {

	var word string
	in := bufio.NewReader(os.Stdin)

	// Print results on  stdout
	fmt.Println("Press enter to add your word to a dictionary, backspace to erase and escape to exit.\r")

loop:
	for {
		// Read rune from stdin
		r, _, err := in.ReadRune()
		if err != nil {
			panic(err)
		}

		switch r {
		// escape
		case 27:
			break loop
		// enter
		case 13:
			root.insert(word)
			fmt.Printf("\x1b[2K\r%s\n\x1b[2K\r%v\x1b[1A\x1b[%dG", word, root.search(word), len(word)+1)
			word = ""
		// backspace
		case 127:
			word = ""
			fmt.Print("\x1b[2K\r")
		default:
			word += string(r)
			fmt.Printf("\x1b[2K\r%s\n\x1b[2K\r%v\x1b[1A\x1b[%dG", word, root.search(word), len(word)+1)
		}
	}
}

func main() {
	root := newNode()

	// Set terminal into raw mode
	state, err := term.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := term.Restore(0, state); err != nil {
			panic(err)
		}
	}()

	handleInput(root)
}
