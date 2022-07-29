package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input struct {
		MessageRef string
	}
	err := json.NewDecoder(strings.NewReader(os.Args[1])).Decode(&input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	react := struct {
		MessageRef string
		Content    string
		Action     string
	}{
		MessageRef: input.MessageRef,
		Content:    "+1",
		Action:     "react",
	}

	reply := struct {
		MessageRef string
		Content    string
		Action     string
	}{
		MessageRef: input.MessageRef,
		Content:    ":+1:",
		Action:     "reply",
	}

	b, err := json.Marshal(&react)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b2, err := json.Marshal(&reply)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("#!#%s\n#!#%s", string(b), string(b2))
}
