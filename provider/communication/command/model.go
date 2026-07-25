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
		fmt.Println(that.Stdout)
	}

	if that.Stderr != "" {
		fmt.Println(that.Stderr)
	}
}

func (that *Result) Log() {

	if that.Stdout != "" {
		log.Println(that.Stdout)
	}

	if that.Stderr != "" {
		log.Println(that.Stderr)
	}
}
