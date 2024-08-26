package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

func main() {
	text := ReadingStdin()
	data := make(map[[2]string][]string)

	for i := 0; i < len(text)-2; i++ {
		key := [2]string{text[i], text[i+1]}
		data[key] = append(data[key], text[i+2])
	}
	fmt.Print(text[0], " ", text[1])
	MarkovChainAlgorithm(text[0], text[1], data, 2)
}

func ReadingStdin() []string {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error accessing stdin:", err)
		os.Exit(1)
	}
	if fileInfo.Mode()&os.ModeNamedPipe == 0 {
		fmt.Fprintln(os.Stderr, "Error: no input text")
		os.Exit(1)
	}

	buf := new(strings.Builder)
	io.Copy(buf, os.Stdin)
	input := buf.String()
	return strings.Fields(input)
}

func MarkovChainAlgorithm(w1, w2 string, data map[[2]string][]string, counter int) {
	prefix := [2]string{w1, w2}
	if len(data[prefix]) == 0 || counter >= 100 {
		fmt.Println()
		return
	}
	w3 := data[prefix][rand.Intn(len(data[prefix]))]
	fmt.Print(" ", w3)
	MarkovChainAlgorithm(w2, w3, data, counter+1)
}
