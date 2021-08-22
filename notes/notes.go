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

func ListNote() {
	notes := loadNotes()
	if len(notes) > 0 {
		fmt.Println("Your Notes")
		for _, note := range notes {
			fmt.Printf("Title: %s\n", note.Title)
			for i, body := range note.Body {
				fmt.Printf("%d. %s\n", i + 1, body)
			}
		}
	} else {
		fmt.Println("Notes are empty")
	}
}

func ReadNote(title string) {
	notes := loadNotes()
	found := findNote(notes, title)
	if found == -1 {
		fmt.Println("Note not found")
	} else {
		body := notes[found].Body
		fmt.Printf("Title: %s\n", notes[found].Title)
		for i, s := range body {
			fmt.Printf("%d. %s\n", i + 1, s)
		}
	}
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

