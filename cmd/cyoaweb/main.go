package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/philip314/cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "JSON file of story")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		exit("Error opening file")
	}

	story, err := cyoa.CreateStory(file)
	if err != nil {
		exit(fmt.Sprintf("Error creating story from JSON file, %s", err))
	}

	fmt.Println(story)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
