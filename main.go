package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

var (
	numberOfWords int
	startPrefix   string
)

func init() {
	flag.IntVar(&numberOfWords, "w", 100, "Number of maximum words")
	flag.StringVar(&startPrefix, "p", "Chapter 1", "Starting prefix") // p can be > l not reverse
	flag.Parse()
}

func main() {
	if numberOfWords < 0 || numberOfWords > 10000 {
		fmt.Fprintln(os.Stderr, "Error: invalid amount of words.")
		os.Exit(1)
	}

	text := ReadingStdin()
	data := make(map[[2]string][]string)

	if startPrefix != "Chapter 1" && !PrefixInText(text) {
		fmt.Fprintln(os.Stderr, "Error: the given prefix is not in the text.")
		os.Exit(1)
	}

	for i := 0; i < len(text)-2; i++ {
		key := [2]string{text[i], text[i+1]}
		data[key] = append(data[key], text[i+2])
	}
	fmt.Print(startPrefix)

	MarkovChainAlgorithm(strings.Split(startPrefix, " ")[0], strings.Split(startPrefix, " ")[1], data, 2)
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
	if len(data[prefix]) == 0 || counter >= numberOfWords {
		fmt.Println()
		return
	}
	w3 := data[prefix][rand.Intn(len(data[prefix]))]
	fmt.Print(" ", w3)
	MarkovChainAlgorithm(w2, w3, data, counter+1)
}

func PrefixInText(text []string) bool {
	inText := false
	for _, word := range text {
		counterPrefix := 0
		for _, prefix := range strings.Split(startPrefix, " ") {
			if word == prefix {
				counterPrefix++
			}
		}
		if counterPrefix == len(strings.Split(startPrefix, " "))-1 {
			inText = true
			break
		}
	}
	return inText
}
