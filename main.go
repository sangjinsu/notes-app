package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type StringArray []string

func (arr *StringArray) String() string {
	return fmt.Sprintf("%v", *arr)
}

func (arr *StringArray) Set(s string) error {
	*arr = strings.Split(s, ",")
	return nil
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Command is needed")
		os.Exit(1)
	}

	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	removeCommand := flag.NewFlagSet("remove", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	readCommand := flag.NewFlagSet("read", flag.ExitOnError)

	addTitleFlag := addCommand.String("title", "", "노트 이름을 입력합니다")
	addBodyFlag := addCommand.String("body", "", "노트 내용을 입력합니다")
	//listAllFlag := listCommand.Bool("all", false, "노트 전체를 출력합니다")
	//removeTitleFlag = removeCommand.String("title", "", "노트 이름을 입력합니다")
	//readTitleFlag = readCommand.String("title", "", "노트 이름을 입력합니다")

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
		if *addTitleFlag == "" || *addBodyFlag == "" {
			addCommand.PrintDefaults()
			os.Exit(1)
		}

		// 노트 더하기
	}
}
