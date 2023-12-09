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
	lookup, seeds, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range lookup {
		fmt.Println(key, len(value))
	}

	fmt.Println("========== SOLUTIONS ==========")
	fmt.Println("Part 1: ", len(seeds))
}

const (
	seedLineId string = "seeds: "

	seedToSoilMapLineId            string = "seed-to-soil map:"
	soilToFertilizerMapLineId      string = "soil-to-fertilizer map:"
	fertilizerToWaterMapLineId     string = "fertilizer-to-water map:"
	waterToLightMapLineId          string = "water-to-light map:"
	lightToTemperatureMapLineId    string = "light-to-temperature map:"
	temperatureToHumidityMapLineId string = "temperature-to-humidity map:"
	humidityToLocationMapLineId    string = "humidity-to-location map:"
)

type mapItem struct {
	Beginning int
	Middle    int
	End       int
}

type mapLookup map[string][]mapItem

func readFile() (mapLookup, []int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return mapLookup{}, []int{}, err
	}
	defer file.Close()

	lookup := mapLookup{
		seedToSoilMapLineId:            []mapItem{},
		soilToFertilizerMapLineId:      []mapItem{},
		fertilizerToWaterMapLineId:     []mapItem{},
		waterToLightMapLineId:          []mapItem{},
		lightToTemperatureMapLineId:    []mapItem{},
		temperatureToHumidityMapLineId: []mapItem{},
		humidityToLocationMapLineId:    []mapItem{},
	}
	var target string
	seeds := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, seedLineId) {
			seedIds := strings.Split(strings.Replace(line, seedLineId, "", 1), " ")
			for _, seedId := range seedIds {
				parsedSeedId, _ := strconv.Atoi(seedId)
				seeds = append(seeds, parsedSeedId)
			}
			continue
		}

		target = processLine(lookup, target, line)
	}
	return lookup, seeds, scanner.Err()
}

func processLine(lookup mapLookup, currentTargetId string, line string) string {
	if len(line) == 0 {
		return ""
	}

	for lineId := range lookup {
		if strings.Contains(line, lineId) {
			return lineId
		}
	}

	splitLine := strings.Split(line, " ")
	beginning, _ := strconv.Atoi(splitLine[0])
	middle, _ := strconv.Atoi(splitLine[1])
	end, _ := strconv.Atoi(splitLine[2])
	lookup[currentTargetId] = append(lookup[currentTargetId], mapItem{
		Beginning: beginning,
		Middle:    middle,
		End:       end,
	})

	return currentTargetId
}
