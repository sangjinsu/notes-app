package main

import (
	"encoding/json"
	"os"
)

type Notes []Note

type Note struct {
	Title string   `json:"title"`
	Body  []string `json:"body"`
}

func SaveNotes(notes Notes) {
	f, err := os.OpenFile("notes.json", os.O_WRONLY, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			f, _ = os.Create("notes.json")
		} else {
			panic(err)
		}
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	encoder := json.NewEncoder(f)
	err = encoder.Encode(notes)
	if err != nil {
		panic(err)
	}
}

func LoadNotes() Notes {
	f, err := os.OpenFile("notes.json", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	encoder := json.NewDecoder(f)

	var notes Notes
	err = encoder.Decode(&notes)
	if err != nil {
		panic(err)
	}

	return notes
}

