package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	field := flag.Int("f", 0, "field flag")
	delimiter := flag.String("d", "\t", "delimiter flag")

	flag.Parse()

	var file *os.File
	var err error
	file, err = os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if *field > 0 {
		result := cutField(file, *field, *delimiter)
		fmt.Println(result)
	} else {
		fmt.Println("invalid field value")
	}

}

func cutField(file *os.File, field int, delimeter string) string {

	var result []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		cols := strings.Split(row, delimeter)
		result = append(result, cols[field-1])
	}

	return strings.Join(result, "\n")
}
