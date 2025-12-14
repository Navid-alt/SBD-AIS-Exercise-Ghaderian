package main

import (
	"bufio"
	"exc9/mapred"
	"fmt"
	"os"
	"sort"
)

// Main function
func main() {
	// todo read file
	text, err := readLines("res/meditations.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("Read %d lines from 'res/meditations.txt'. Running MapReduce...\n", len(text))
	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)
	// todo print your result to stdout
	printTopResults(results, 20)

}
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func printTopResults(results map[string]int, n int) {
	type kv struct {
		Key   string
		Value int
	}
	var sortedResults []kv
	for k, v := range results {
		sortedResults = append(sortedResults, kv{k, v})
	}
	sort.Slice(sortedResults, func(i, j int) bool {
		return sortedResults[i].Value > sortedResults[j].Value
	})

	fmt.Println("\n--- Top Word Frequencies ---")
	for i := 0; i < n && i < len(sortedResults); i++ {
		fmt.Printf("%-15s : %d\n", sortedResults[i].Key, sortedResults[i].Value)
	}
}
