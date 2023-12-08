package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := readFile()
	if err != nil {
		log.Fatal(err)
	}



	fmt.Println("========== SOLUTIONS ==========")
	fmt.Println("Part 1:", -1)
}

func readFile() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
