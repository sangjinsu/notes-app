package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Notes []Note

type Note struct {
	Title string   `json:"title"`
	Body  []string `json:"body"`
}

func AddNotes(title, body string) {
	notes := loadNotes()
	found := findNote(notes, title)

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

func RemoveNote(title string) {
	notes := loadNotes()
	found := findNote(notes, title)
	if found == -1 {
		fmt.Println("Note not found")
	} else {
		notes = append(notes[:found], notes[found+1:]...)
	}
	saveNotes(notes)
}

func saveNotes(notes Notes) {
	bytes, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("notes.json", bytes, 0755)
	if err != nil {
		panic(err)
	}
}

func loadNotes() Notes {
	file, err := ioutil.ReadFile("notes.json")
	if err != nil {
		panic(err)
	}
	var notes Notes
	err = json.Unmarshal(file, &notes)
	if err != nil {
		panic(err)
	}
	return notes
}

func findNote(notes Notes, title string) int {
	for i, n := range notes {
		if n.Title == title {
			return i
		}
	}
	return -1
}

func main() {
	note := Note{Title: "hello", Body: []string{"hello2"}}
	var notes Notes
	notes = []Note{note, {Title: "hello2", Body: []string{"hello3"}}}
	saveNotes(notes)

	AddNotes("hello", "bye2")
	RemoveNote("hello")
}
