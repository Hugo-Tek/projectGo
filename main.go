package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	X, Y int
}

type Map struct {
	X int
	Y int
}

type error interface {
	Error() string
}

func firstLine(firstLine []string, entrepot Map) {
	entrepot.X, err = strconv.Atoi(firstLine[0])
	entrepot.Y, err = strconv.Atoi(firstLine[1])
	if err != nil {
		// ... handle error
		panic(err)
	}
	fmt.Println(entrepot.X)
}

func fileParser(lines []string, entrepot Map) {
	firstLine(strings.Split(lines[0], " "), entrepot)
}

// open a file and read it line by line and print it to the console
func main() {
	entrepot := Map{0, 0}
	fileToParse := os.Args[1]
	file, err := os.Open(fileToParse)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fileParser(lines, entrepot)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
