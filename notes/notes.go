package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Notes []Note

type Note struct {
	Title string   `json:"title"`
	Body  []string `json:"body"`
}

func AddNotes(title, body string) {
	notes := loadNotes()

	found := -1
	for i, n := range notes {
		if n.Title == title {
			found = i
			break
		}
	}

	if found == -1 {
		notes = append(notes, Note{Title: title, Body: []string{body}})
		fmt.Println("New note added")
	} else {
		fmt.Println("Note is already existed")
		notes[found].Body = append(notes[found].Body, body)
		fmt.Println("body added")
	}
	saveNotes(notes)
}

func saveNotes(notes Notes) {
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

func loadNotes() Notes {
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

	notes := new(Notes)
	err = encoder.Decode(notes)
	if err != nil {
		panic(err)
	}
	return *notes
}

func main() {
	note := Note{Title: "hello", Body: []string{"hello2"}}
	var notes Notes
	notes = []Note{note, {Title: "hello2", Body: []string{"hello3"}}}
	saveNotes(notes)

	AddNotes("hello", "bye2")
}
