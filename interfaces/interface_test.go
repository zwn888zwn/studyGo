package interfaces

import (
	"fmt"
	"os"
	"testing"
)

func TestInterface1(t *testing.T) {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello world")
	fmt.Fprintln(os.Stdout, UpperString("hello world"))
}
