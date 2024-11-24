package fields

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func CutField(file *os.File, parsedFields []int, delimeter string) string {
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

func ParseFields(fields string) []int {
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
