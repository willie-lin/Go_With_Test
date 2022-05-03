package mocking

import (
	"fmt"
	"io"
)

func Countdown(out io.Writer) {
	fmt.Print(out, "3")
}
