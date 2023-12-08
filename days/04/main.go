package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========== SOLUTIONS ==========")
	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
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

	return totalPoints
}

func part2(lines []string) int {
	cards := make(map[int]int, len(lines))
	for i := 1; i <= len(lines); i++ {
		cards[i] = 1
	}

	for _, card := range lines {
		id := strings.Replace(strings.Split(card, ":")[0], "Card ", "", 1)
		cId, _ := strconv.ParseInt(strings.TrimSpace(id), 0, 64)
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

		wins := 0
		for _, drawingNumber := range drawingNumbers {
			if _, ok := winningNumberLookup[strings.TrimSpace(drawingNumber)]; ok {
				wins++
				continue
			}
		}

		cardId := int(cId)
		copiesOfCurrentCard := cards[cardId]
		for i := 1; i <= wins; i++ {
			cards[cardId+i] += copiesOfCurrentCard
		}
	}

	totalPoints := 0
	for _, numCards := range cards {
		totalPoints += numCards
	}

	return totalPoints
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
