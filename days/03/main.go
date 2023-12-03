package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

type number struct {
	Value int
	Line  int
	Start int
	End   int
}

type symbol struct {
	Value  string
	Line   int
	Column int
}

func main() {
	lines, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	numbers := []number{}
	symbols := []symbol{}
	for lineIdx, line := range lines {
		startIdx := -1
		accumulator := 0
		for runeIdx, rune := range line {
			if unicode.IsDigit(rune) {
				if startIdx == -1 {
					startIdx = runeIdx
				}
				number, _ := strconv.Atoi(string(rune))
				accumulator = accumulator*10 + number
				continue
			}
			if startIdx > -1 {
				numbers = append(numbers, number{
					Value: accumulator,
					Line:  lineIdx,
					Start: startIdx,
					End:   runeIdx - 1,
				})
				accumulator = 0
				startIdx = -1
			}

			if rune == '.' {
				continue
			}

			symbols = append(symbols, symbol{
				Value:  string(rune),
				Line:   lineIdx,
				Column: runeIdx,
			})
		}
		if startIdx > -1 {
			numbers = append(numbers, number{
				Value: accumulator,
				Line:  lineIdx,
				Start: startIdx,
				End:   len(line),
			})
		}
	}

	partItems := []number{}
	gearParts := map[string][]number{}
	for _, item := range numbers {
		if symbol, found := hasAdjacentSymbol(symbols, item); found {
			partItems = append(partItems, item)
			if symbol.Value == "*" {
				symbolId := fmt.Sprintf("%d,%d", symbol.Line, symbol.Column)
				gearParts[symbolId] = append(gearParts[symbolId], item)
			}
		}
	}

	sumOfPartNumbers := 0
	for _, item := range partItems {
		sumOfPartNumbers += item.Value
	}

	gearRatioTotals := 0
	for _, gearPart := range gearParts {
		if len(gearPart) != 2 {
			continue
		}
		gearRatioTotals += gearPart[0].Value * gearPart[1].Value
	}

	fmt.Println("sum of part numbers:", sumOfPartNumbers)
	fmt.Println("gear ratio totals:", gearRatioTotals)
}

func hasAdjacentSymbol(symbols []symbol, num number) (symbol, bool) {
	for _, symbol := range symbols {
		lineDelta := int(math.Abs(float64(num.Line - symbol.Line)))
		startDelta := int(math.Abs(float64(num.Start - symbol.Column)))
		endDelta := int(math.Abs(float64(num.End - symbol.Column)))
		if lineDelta <= 1 && (startDelta <= 1 || endDelta <= 1) {
			return symbol, true
		}
	}
	return symbol{}, false
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
