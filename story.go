package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// Construct a story object from a JSON file
func CreateStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}
