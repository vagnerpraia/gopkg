package gpcommand

import "fmt"

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
