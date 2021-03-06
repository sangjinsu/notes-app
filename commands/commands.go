package commands

import (
	"flag"
	"fmt"
	"github.com/notes/handleerror"
	"github.com/notes/notes"
	"os"
)

var err error

func Init() {
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	removeCommand := flag.NewFlagSet("remove", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	readCommand := flag.NewFlagSet("read", flag.ExitOnError)

	var addTitleFlag, addBodyFlag, removeTitleFlag, readTitleFlag string
	var listAllFlag bool

	addCommand.StringVar(&addTitleFlag, "title", "", "노트 이름을 입력합니다")
	addCommand.StringVar(&addBodyFlag, "body", "", "노트 내용을 입력합니다")
	listCommand.BoolVar(&listAllFlag, "all", false, "노트 전체를 출력합니다")
	removeCommand.StringVar(&removeTitleFlag, "title", "", "노트 이름을 입력합니다")
	readCommand.StringVar(&readTitleFlag, "title", "", "노트 이름을 입력합니다")

	switch command := os.Args[1]; command {
	case "add":
		err = addCommand.Parse(os.Args[2:])
		if err != nil {
			panic(handleerror.MakeErr("명령을 수행할 수 없습니다: ", err))
		}
	case "remove":
		err = removeCommand.Parse(os.Args[2:])
		if err != nil {
			panic(handleerror.MakeErr("명령을 수행할 수 없습니다: ", err))
		}
	case "list":
		err = listCommand.Parse(os.Args[2:])
		if err != nil {
			panic(handleerror.MakeErr("명령을 수행할 수 없습니다: ", err))
		}
	case "read":
		err = readCommand.Parse(os.Args[2:])
		if err != nil {
			panic(handleerror.MakeErr("명령을 수행할 수 없습니다: ", err))
		}
	default:
		fmt.Println("Usage of add:")
		addCommand.PrintDefaults()
		fmt.Println("Usage of remove:")
		removeCommand.PrintDefaults()
		fmt.Println("Usage of list:")
		listCommand.PrintDefaults()
		fmt.Println("Usage of read:")
		readCommand.PrintDefaults()
		os.Exit(1)
	}

	if addCommand.Parsed() {
		if addTitleFlag == "" || addBodyFlag == "" {
			addCommand.PrintDefaults()
			os.Exit(1)
		}
		notes.AddNotes(addTitleFlag, addBodyFlag)
	}

	if removeCommand.Parsed() {
		if removeTitleFlag == "" {
			removeCommand.PrintDefaults()
			os.Exit(1)
		}
		notes.RemoveNote(removeTitleFlag)
	}

	if listCommand.Parsed() {
		if listAllFlag {
			notes.ListAllNote()
		} else {
			notes.ListTitleNote()
		}
	}

	if readCommand.Parsed() {
		if readTitleFlag == "" {
			removeCommand.PrintDefaults()
			os.Exit(1)
		}
		notes.ReadNote(readTitleFlag)
	}
}
