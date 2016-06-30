package main

import "fmt"
import "os"
import "strings"

func EjemploCriptografia() {
	/*
	  fmt.Println("Test!")
	  text := os.Args[1]
	  fmt.Println(text)
	  fmt.Println(len(text))
	  for i:=0; i < len(text); i++ {
	    fmt.Println(string(text[i]))
	  }
	*/

	if len(os.Args) > 1 {
		text := os.Args[1:]
		fmt.Println(process(strings.ToLower(strings.Join(text, " "))))
	} else {
		fmt.Println("Must enter at least one parameter")
	}

}

func process(message string) string {
	var result string
	for i := 0; i < len(message); i++ {
		result += encrypt(string(message[i]))
	}
	return result
}

func encrypt(letter string) string {

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	if letter == " " {
		return "-"
	}
	foundIndex := strings.Index(alphabet, letter)
	newIndex := (foundIndex + 1) % len(alphabet)
	if foundIndex == -1 {
		return string("?")
	}
	return string(alphabet[newIndex])

}
