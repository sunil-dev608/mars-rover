package parser

import (
	"fmt"
	"regexp"
	"strconv"
)

type ParsedData struct {
	X           int
	Y           int
	Orientation rune
	Command     string
}

type Parser interface {
	ParseRobotString(input string) (*ParsedData, error)
}

type parserImpl struct {
}

func NewParser() Parser {
	return &parserImpl{}
}

func (w *parserImpl) ParseRobotString(input string) (*ParsedData, error) {
	// Define the regular expression to match the format
	re := regexp.MustCompile(`^\((\d+),\s*(\d+),\s*(\w+)\)\s+(\w+)$`)

	// Find submatches
	matches := re.FindStringSubmatch(input)
	if matches == nil || len(matches) != 5 {
		return nil, fmt.Errorf("input does not match the required format")
	}

	// Parse the integers
	a, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}

	b, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	// Fill the struct with parsed values
	return &ParsedData{
		X:           a,
		Y:           b,
		Orientation: (rune)(matches[3][0]),
		Command:     matches[4],
	}, nil
}
