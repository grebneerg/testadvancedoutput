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
		Channel    struct {
			ID string
		}
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

	bookmark := struct {
		Title     string
		Content   string
		Action    string
		ChannelID string
	}{
		Title:     "test",
		Content:   "https://jackg.dev",
		Action:    "bookmark",
		ChannelID: input.Channel.ID,
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
	b3, err := json.Marshal(&bookmark)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("#!#%s\n#!#%s\n#!#%s", string(b), string(b2), string(b3))
}
