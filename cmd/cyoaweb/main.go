package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/philip314/cyoa"
)

func main() {
	portFlag := flag.Int("port", 8080, "Port that application runs on")
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

	handler := cyoa.StoryHandler(story)
	port := fmt.Sprintf(":%d", *portFlag)

	log.Fatal(http.ListenAndServe(port, handler))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
