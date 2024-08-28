---

# markov-chain

## Overview

The **markov-chain** project involves building a text generator using the Markov Chain algorithm. This algorithm is commonly used in applications like predictive text on smartphones, where it suggests the next most probable word based on previous input.

## Learning Objectives

- Understand and implement algorithms
- Gain proficiency in I/O and file handling in Go
- Learn the importance of data structures in software design

## Project Details

In this project, you will develop a Markov Chain text generator that reads an English text from `stdin`, processes it, and generates new text based on the frequency of fixed-length phrases. The algorithm creates random sequences of words that resemble natural language by analyzing the input text and predicting the most likely next word based on prefixes of a specified length.

## Features

- Generates random text using the Markov Chain algorithm
- Supports setting custom prefix length and starting prefix
- Limits the output text to a maximum number of words
- Handles error cases, such as invalid input or missing text

## Usage

### Build the Program

Compile the program by running the following command in the project's root directory:

```bash
$ go build -o markovchain .
```

### Run the Program

To generate text based on the Markov Chain algorithm, provide input through `stdin` and use the following options:

```bash
$ cat input.txt | ./markovchain [-w <N>] [-p <S>] [-l <N>]
```

### Options

- `-w <N>`: Set the maximum number of words to generate (default: 100)
- `-p <S>`: Specify the starting prefix (default: first N words of the input text)
- `-l <N>`: Set the prefix length (default: 2, max: 5)
- `--help`: Show usage information

### Examples

Generate 100 words starting with the prefix "Chapter 1":

```bash
$ cat input.txt | ./markovchain -p "Chapter 1" -w 100
```

Generate text with a 3-word prefix:

```bash
$ cat input.txt | ./markovchain -l 3 -w 50
```

## Guidelines

- Focus on designing efficient data structures to handle large input sizes (e.g., a book with 100,000 words).
- Your program should produce natural-sounding text by considering punctuation as part of the words.
- Ensure the program handles errors gracefully and exits with an appropriate message.

## Support

If you encounter issues, start by testing your program with example inputs from the project. Verify that your output matches the expected results. If you're stuck, review the project description and try debugging with smaller test cases.

## Author

Created by Alimukhamed Tlekbai, Team Lead at Doodocs.kz

Contact:
- Email: tlekbai@doodocs.kz
- [GitHub](https://github.com/your-github)
- [LinkedIn](https://linkedin.com/in/your-linkedin)

---