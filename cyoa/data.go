package main

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

func JsonStory(r io.Reader) (Story, error) {
	var story Story
	err := json.NewDecoder(r).Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}
