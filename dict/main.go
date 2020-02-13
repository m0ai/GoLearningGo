package main

import (
	"fmt"

	"github.com/m0ai/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)

	word, err := dictionary.Search(baseWord)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
