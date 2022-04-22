package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	filemap := make(map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		log.Fatal("You must pass at least 1 file name")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filemap)
			f.Close()
		}
	}
	for k := range filemap {
		fmt.Printf("File with duplicate lines = %s\n", k)
	}
}

func countLines(f *os.File, counts map[string]int, filemap map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		_, ok := counts[line]
		if ok {
			filemap[f.Name()] = true
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
