package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	text := ReadingStdin()

	// // map_of_pas := make(map[string][]string)
	fmt.Println(len(text))
}

func ReadingStdin() []string {
	file_info, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error accessing stdin:", err)
		os.Exit(1)
	}
	if file_info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprintln(os.Stderr, "Error: no input text")
		os.Exit(1)
	}

	buf := new(strings.Builder)
	io.Copy(buf, os.Stdin)
	input := buf.String()
	return strings.Fields(input)
}
