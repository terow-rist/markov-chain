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
	flag.StringVar(&startPrefix, "p", "", "Starting prefix")
	flag.IntVar(&prefixLength, "l", 2, "Prefix length")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Markov Chain text generator.")
		fmt.Fprintln(os.Stderr, "\nUsage:\n  markovchain [-w <N>] [-p <S>] [-l <N>]\n  markovchain --help")
		fmt.Fprintln(os.Stderr, "\nOptions:\n  --help  Show this screen.\n  -w N    Number of maximum words\n  -p S    Starting prefix\n  -l N    Prefix length")
	}

	flag.Parse()
}

func main() {
	text := ReadingStdin()
	if len(text) < 2 {
		fmt.Fprintln(os.Stderr, "Error: not enough words to generate prefix")
		os.Exit(1)
	}

	if prefixLength < 1 || prefixLength > 5 {
		fmt.Fprintln(os.Stderr, "Error: length of prefix must be in between 1 and 5")
		os.Exit(1)
	}

	if numberOfWords < 1 || numberOfWords > 10000 {
		fmt.Fprintln(os.Stderr, "Error: number of words must be in between 1 and 10000")
		os.Exit(1)
	}

	if len(startPrefix) == 0 {
		startPrefix = strings.Join(text[:prefixLength], " ")
	} else if len(strings.Fields(startPrefix)) != prefixLength {
		fmt.Fprintf(os.Stderr, "Error: prefix must be '%d' words\n", prefixLength)
		os.Exit(1)
	}

	if !PrefixInText(text) {
		fmt.Fprintln(os.Stderr, "Error: prefix not found")
		os.Exit(1)
	}

	data := make(map[string][]string)
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
		if sliceToString(text[i:i+len(startPrefixWords)]) == startPrefix {
			return true
		}
	}
	return false
}

func sliceToString(slice []string) string {
	return strings.Join(slice, " ")
}
