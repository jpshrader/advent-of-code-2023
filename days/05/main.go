package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lookup, seeds, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("========== SOLUTIONS ==========")
	fmt.Println("Part 1: ", part1(lookup, seeds))
}

func part1(lookup mapLookup, seeds []int) int {
	locations := []int{}
	for _, seed := range seeds {
		soil := findMapMatches(lookup, seedToSoilMapLineId, seed)
		fertilizer := findMapMatches(lookup, soilToFertilizerMapLineId, soil)
		water := findMapMatches(lookup, fertilizerToWaterMapLineId, fertilizer)
		light := findMapMatches(lookup, waterToLightMapLineId, water)
		temperature := findMapMatches(lookup, lightToTemperatureMapLineId, light)
		humidity := findMapMatches(lookup, temperatureToHumidityMapLineId, temperature)
		location := findMapMatches(lookup, humidityToLocationMapLineId, humidity)

		locations = append(locations, location)
	}
	return slices.Min(locations)
}

func findMapMatches(lookup mapLookup, currentTargetId string, source int) int {
	for _, mapItem := range lookup[currentTargetId] {
		match := mapItem.FindMatch(source)
		if match != -1 {
			return match
		}
	}
	return source
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
	DestinationRange int
	SourceRangeStart int
	RangeLength      int
}

func (mi mapItem) FindMatch(source int) int {
	if source >= mi.SourceRangeStart && source < mi.SourceRangeStart+mi.RangeLength {
		sourceOffset := source - mi.SourceRangeStart
		return mi.DestinationRange + sourceOffset
	}
	return -1
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
	destinationRange, _ := strconv.Atoi(splitLine[0])
	sourceRangeStart, _ := strconv.Atoi(splitLine[1])
	rangeLength, _ := strconv.Atoi(splitLine[2])
	lookup[currentTargetId] = append(lookup[currentTargetId], mapItem{
		DestinationRange: destinationRange,
		SourceRangeStart: sourceRangeStart,
		RangeLength:      rangeLength,
	})

	return currentTargetId
}
