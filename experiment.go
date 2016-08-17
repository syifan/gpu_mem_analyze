package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type DataPoint struct {
	Iteration int
	Accuracy  float64
	TimeInSec float64
}

type Experiment struct {
	StartTime time.Time
	BatchSize int
	Points    []DataPoint
}

func LoadExperimentFromFile(path string) (experiment Experiment) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	lineNumber := 0
	var startTime time.Time
	currentIteration := 0
	for {
		lineNumber++
		line, eof := retrieveLine(reader)
		if eof {
			break
		}

		if lineNumber == 1 {
			startTime = parseForTime(line)
			experiment.StartTime = startTime
			// fmt.Println(startTime)
		} else {
			iteration := tryParseIteration(line)
			if iteration != 0 {
				currentIteration = iteration
			}

			dataPoint := tryToParseDataPoint(line, startTime)
			if dataPoint != nil {
				dataPoint.Iteration = currentIteration
				experiment.Points = append(experiment.Points, *dataPoint)
			}
		}

	}

	fmt.Println(experiment)
	return
}

func retrieveLine(reader *bufio.Reader) (line string, eof bool) {
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		eof = true
		return
	} else if err != nil {
		log.Fatal(err)
	}

	line = strings.TrimSuffix(line, "\n")
	return
}

func parseForTime(line string) (startTime time.Time) {
	pattern := regexp.MustCompile(`[0-9]+:[0-9]+:[0-9]+\.[0-9]+`)
	subMatch := string(pattern.FindSubmatch([]byte(line))[0])

	layout := "15:04:05.000000"
	startTime, err := time.Parse(layout, subMatch)
	if err != nil {
		log.Panic(err)
	}

	return
}

func tryParseIteration(line string) (iteration int) {
	pattern := regexp.MustCompile(`Iteration ([0-9]+)`)
	match := pattern.FindSubmatch([]byte(line))

	if len(match) == 0 {
		return 0
	} else {
		iteration, _ = strconv.Atoi(string(match[1]))
		return
	}
}

func tryToParseDataPoint(line string, startTime time.Time) (dataPoint *DataPoint) {
	accuracyPattern := regexp.MustCompile(`accuracy = ([0-9\.]+)`)
	match := accuracyPattern.FindSubmatch([]byte(line))

	if len(match) == 0 {
		return nil
	}

	accuracy, _ := strconv.ParseFloat(string(match[1]), 64)
	dataPoint = new(DataPoint)
	dataPoint.Accuracy = accuracy

	time := parseForTime(line)
	timeDiff := 1.0 * float64(time.Sub(startTime)) / 1e9
	dataPoint.TimeInSec = timeDiff

	return
}
