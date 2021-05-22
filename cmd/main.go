package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rogercoll/autovpn"
)

// StdoutLogger represents the stdout logger callback
type StdoutLogger func(text string)

// Log logs the given string to stdout logger
func (lc StdoutLogger) Log(text string) {
	lc(text)
}

func main() {
	profileFile := os.Args[1]
	var logger StdoutLogger = func(text string) {
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			fmt.Println("Library check >>", line)
		}
	}
	conf, err := autovpn.NewConfig("a", "b", "c", logger, nil)
	if err != nil {
		panic(err)
	}
	conf.Start()
}
