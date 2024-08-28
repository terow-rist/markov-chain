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
	prefixLength  int
)

func init() {
	flag.IntVar(&numberOfWords, "w", 100, "Number of maximum words")
	flag.StringVar(&startPrefix, "p", "Chapter 1", "Starting prefix")
	flag.IntVar(&prefixLength, "l", 2, "Prefix length")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Markov Chain text generator.")
		fmt.Fprintln(os.Stderr, "\nUsage:\n  markovchain [-w <N>] [-p <S>] [-l <N>]\n  markovchain --help")
		fmt.Fprintln(os.Stderr, "\nOptions:\n  --help  Show this screen.\n  -w N    Number of maximum words\n  -p S    Starting prefix\n  -l N    Prefix length")
	}

	flag.Parse()
}

func main() {
	if numberOfWords < 0 || numberOfWords > 10000 {
		fmt.Fprintln(os.Stderr, "Error: invalid amount of words.")
		os.Exit(1)
	} else if len(strings.Split(startPrefix, " ")) != prefixLength || prefixLength < 0 || prefixLength > 5 {
		fmt.Fprintln(os.Stderr, "Error: incorrect input for prefix length.")
		os.Exit(1)
	}

	text := ReadingStdin()
	data := make(map[string][]string)

	if startPrefix != "Chapter 1" && !PrefixInText(text) {
		fmt.Fprintln(os.Stderr, "Error: the given prefix is not in the text.")
		os.Exit(1)
	}

	for i := 0; i < len(text)-prefixLength; i++ {
		key := sliceToString(text[i : i+prefixLength])
		data[key] = append(data[key], text[i+prefixLength])
	}
	fmt.Print(startPrefix)

	MarkovChainAlgorithm(strings.Fields(startPrefix), data, prefixLength)
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

func MarkovChainAlgorithm(prefix []string, data map[string][]string, counter int) {
	if len(data[sliceToString(prefix)]) == 0 || counter >= numberOfWords {
		fmt.Println()
		return
	}
	w3 := data[sliceToString(prefix)][rand.Intn(len(data[sliceToString(prefix)]))]
	fmt.Print(" ", w3)

	newPrefix := append(prefix[1:], w3)

	MarkovChainAlgorithm(newPrefix, data, counter+1)
}

func PrefixInText(text []string) bool {
	startPrefixWords := strings.Fields(startPrefix)

	for i := 0; i <= len(text)-len(startPrefixWords); i++ {
		matches := true
		for j := 0; j < len(startPrefixWords); j++ {
			if text[i+j] != startPrefixWords[j] {
				matches = false
				break
			}
		}

		if matches {
			return true
		}
	}

	return false
}

func sliceToString(slice []string) string {
	return strings.Join(slice, " ")
}
