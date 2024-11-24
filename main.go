package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "0", "fields flag")
	delimiter := flag.String("d", "\t", "delimiter flag")
	flag.Parse()
	args := flag.Args()

	var file *os.File

	if len(args) == 0 || args[0] == "-" {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	parsedFields := parseFields(*fields)
	result := cutField(file, parsedFields, *delimiter)
	fmt.Println(result)
}

func cutField(file *os.File, parsedFields []int, delimeter string) string {
	var result []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line string
		row := scanner.Text()
		cols := strings.Split(row, delimeter)
		for i, field := range parsedFields {
			line += cols[field-1]
			if i != len(parsedFields)-1 {
				line += delimeter
			}
		}
		result = append(result, line)
	}
	return strings.Join(result, "\n")
}

func parseFields(fields string) []int {
	normalized := strings.ReplaceAll(fields, ",", " ")
	parts := strings.Fields(normalized)

	var result []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatal(err)
		} else {
			result = append(result, num)
		}
	}
	return result
}
