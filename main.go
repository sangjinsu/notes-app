package main

import (
	"fmt"
	"github.com/notes/commands"
	"log"
	"os"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		log.Fatalln(err)
	} else {
		panic(p)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Command is needed")
		os.Exit(1)
	}

	defer reportPanic()
	commands.Init()
}
