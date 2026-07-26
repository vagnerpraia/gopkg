package gpcommand

import (
	"fmt"
	"log"
)

type Result struct {
	Stdout string
	Stderr string
}

func (that *Result) Print() {

	if that.Stdout != "" {
		fmt.Print(that.Stdout)
	}

	if that.Stderr != "" {
		fmt.Print(that.Stderr)
	}
}

func (that *Result) Log() {

	if that.Stdout != "" {
		log.Print(that.Stdout)
	}

	if that.Stderr != "" {
		log.Print(that.Stderr)
	}
}
