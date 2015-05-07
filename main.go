package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	io.Copy(os.Stdout, os.Stdin)
}
