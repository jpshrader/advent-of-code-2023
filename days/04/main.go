package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	totalPoints := 0
	for _, card := range lines {
		card = strings.Split(card, ":")[1]
		cardNumbers := strings.Split(card, "|")
		winningNumbers := strings.Split(cardNumbers[0], " ")
		drawingNumbers := strings.Split(cardNumbers[1], " ")

		winningNumberLookup := make(map[string]bool, len(winningNumbers))
		for _, winningNumber := range winningNumbers {
			if len(winningNumber) > 0 {
				winningNumberLookup[strings.TrimSpace(winningNumber)] = true
			}
		}

		cardValue := 0
		numberMultipler := 1
		for _, drawingNumber := range drawingNumbers {
			if _, ok := winningNumberLookup[strings.TrimSpace(drawingNumber)]; ok {
				cardValue = numberMultipler
				numberMultipler = numberMultipler * 2
				continue
			}
		}
		totalPoints += cardValue
	}

	fmt.Println("========== SOLUTIONS ==========")
	fmt.Println("Part 1:", totalPoints)
	//fmt.Println("Part 2:", "")
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
