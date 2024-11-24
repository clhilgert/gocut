package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/clhilgert/gocut/cmd/internal/fields"
)

func main() {
	fieldSpec := flag.String("f", "0", "Specify fields to extract (comma or whitespace separated)")
	delimiter := flag.String("d", "\t", "Specify field delimiter (default: tab)")
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

	parsedFields := fields.ParseFields(*fieldSpec)
	result := fields.CutField(file, parsedFields, *delimiter)
	fmt.Println(result)
}
