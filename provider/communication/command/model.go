package gpcommand

import "fmt"

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
