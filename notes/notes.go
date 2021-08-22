package notes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/notes/handleerror"
	"io/ioutil"
	"os"
)

type Notes []Note

type Note struct {
	Title string   `json:"title"`
	Body  []string `json:"body"`
}

var writer = bufio.NewWriter(os.Stdout)

func AddNotes(title, body string) {
	defer writer.Flush()
	notes := loadNotes()
	found := findNote(notes, title)

	if found == -1 {
		notes = append(notes, Note{Title: title, Body: []string{body}})
		fmt.Fprintln(writer, "New note added")
	} else {
		fmt.Fprintln(writer, "Note is already existed")
		notes[found].Body = append(notes[found].Body, body)
		fmt.Fprintln(writer, "body added")
	}
	saveNotes(notes)
}

func RemoveNote(title string) {
	defer writer.Flush()
	notes := loadNotes()
	found := findNote(notes, title)
	if found == -1 {
		fmt.Fprintln(writer, "Note not found")
	} else {
		notes = append(notes[:found], notes[found+1:]...)
		fmt.Fprintln(writer, "Note removed")
	}
	saveNotes(notes)
}

func ListAllNote() {
	defer writer.Flush()
	notes := loadNotes()
	if len(notes) > 0 {
		fmt.Fprintln(writer, "Your Notes")
		for _, note := range notes {
			fmt.Printf("Title: %s\n", note.Title)
			for i, body := range note.Body {
				fmt.Fprintf(writer, "%d. %s\n", i+1, body)
			}
		}
	} else {
		fmt.Fprintln(writer, "Notes are empty")
	}
}

func ListTitleNote() {
	defer writer.Flush()
	notes := loadNotes()
	if len(notes) > 0 {
		fmt.Fprintln(writer, "Your Notes")
		for _, note := range notes {
			fmt.Fprintf(writer, "%s\n", note.Title)
		}
	} else {
		fmt.Fprintln(writer, "Notes are empty")
	}
}

func ReadNote(title string) {
	defer writer.Flush()
	notes := loadNotes()
	found := findNote(notes, title)
	if found == -1 {
		fmt.Fprintln(writer, "Note not found")
	} else {
		body := notes[found].Body
		fmt.Fprintf(writer, "Title: %s\n", notes[found].Title)
		for i, s := range body {
			fmt.Fprintf(writer, "%d. %s\n", i+1, s)
		}
	}
}

func saveNotes(notes Notes) {
	bytes, err := json.Marshal(notes)
	if err != nil {
		panic(handleerror.MakeErr("json 파일로 변환할 수 없습니다", err))
	}
	err = ioutil.WriteFile("notes.json", bytes, 0755)
	if err != nil {
		panic(handleerror.MakeErr("json 파일에 작성할 수 없습니다", err))
	}
}

func loadNotes() Notes {
	var notes Notes
	file, err := ioutil.ReadFile("notes.json")
	if err != nil {
		if os.IsNotExist(err) {
			return notes
		} else {
			panic(handleerror.MakeErr("파일을 읽을 수 없습니다", err))
		}
	}

	err = json.Unmarshal(file, &notes)
	if err != nil {
		panic(handleerror.MakeErr("json 파일을 읽을 수 없습니다", err))
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
