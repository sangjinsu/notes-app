package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Command is needed")
		os.Exit(1)
	}

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	removeCommand := flag.NewFlagSet("remove", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	readCommand := flag.NewFlagSet("read", flag.ExitOnError)

	var addTitleFlag, addBodyFlag, removeTitleFlag, readTitleFlag string
	var listAllFlag bool
	addCommand.StringVar(&addTitleFlag, "title", "", "노트 이름을 입력합니다")
	addCommand.StringVar(&addBodyFlag,"body", "", "노트 내용을 입력합니다")
	listCommand.BoolVar(&listAllFlag, "all", false, "노트 전체를 출력합니다")
	removeCommand.StringVar(&removeTitleFlag, "title", "", "노트 이름을 입력합니다")
	readCommand.StringVar(&readTitleFlag,"title", "", "노트 이름을 입력합니다")

	switch command := os.Args[1]; command {
	case "add":
		addCommand.Parse(os.Args[2:])
	case "remove":
		removeCommand.Parse(os.Args[2:])
	case "list":
		listCommand.Parse(os.Args[2:])
	case "read":
		readCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		fmt.Println("Press ./notes -help")
		os.Exit(1)
	}

	if addCommand.Parsed() {
		if addTitleFlag == "" || addBodyFlag == "" {
			addCommand.PrintDefaults()
			os.Exit(1)
		}

		// 노트 더하기
	}

	if removeCommand.Parsed() {
		if removeTitleFlag == "" {
			removeCommand.PrintDefaults()
			os.Exit(1)
		}
	}

	if listCommand.Parsed() {
		if listAllFlag {

		} else {

		}
	}

	if readCommand.Parsed() {
		if readTitleFlag == "" {
			removeCommand.PrintDefaults()
			os.Exit(1)
		}
	}
}
