package main

import (
	"fmt"
	"github.com/notes/commands"
	"github.com/notes/handleerror"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Command is needed")
		os.Exit(1)
	}

	defer handleerror.ReportPanic()
	commands.Init()
}
