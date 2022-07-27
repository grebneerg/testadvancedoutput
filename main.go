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
	err := json.NewDecoder(strings.NewReader(os.Args[0])).Decode(&input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := struct {
		MessageRef string
		Content    string
	}{
		MessageRef: input.MessageRef,
		Content:    "+1",
	}

	err = json.NewEncoder(os.Stdout).Encode(&output)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
