package keyboardekharab

import (
	"strings"
)

type Keyboard struct {
	durability int
}

func NewKeyboard(durability int) *Keyboard {
	return &Keyboard{
		durability: durability,
	}
}

func (keyboard *Keyboard) Enter(input string) string {
	output := ""
	for _, char := range input {
		if keyboard.durability > 0 {
			if strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 ?!'", string(char)) {
				output += string(char)
			}
			if !strings.Contains("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", string(char)) {
				keyboard.durability--
			}
		}
	}
	return output
}

func main() {
	keyboard := NewKeyboard(2)
	input := "Welcome to CodeCup7"
	output := keyboard.Enter(input)
	println(output)
}

// output: elcome to Cdup7

// example:
//  "Hello, World!"
//  1
//  "Hello, Wrld!"

// "How are you today?"
// 3
// "How are you today?"

// "This is a test!"
// 2
// "This is a test!"
