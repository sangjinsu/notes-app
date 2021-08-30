package handleerror

import (
	"fmt"
	"github.com/notes/chalk"
	"log"
)

func ReportPanic() {
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

func MakeErr(s string, err error) error {
	return fmt.Errorf("%s%s: %s%s", chalk.Red, s, err, chalk.Reset)
}
