package gpcommand

import (
	"fmt"
	"log"
)

type Result struct {
	Command string
	Stdout  string
	Stderr  string
}

func (that *Result) Print() {

	fmt.Printf("command %q executed\n", that.Command)

	if that.Stdout != "" {
		fmt.Println(that.Stdout)
	}

	if that.Stderr != "" {
		fmt.Println(that.Stderr)
	}
}

func (that *Result) Log() {

	log.Printf("command %q executed\n", that.Command)

	if that.Stdout != "" {
		log.Println(that.Stdout)
	}

	if that.Stderr != "" {
		log.Println(that.Stderr)
	}
}
