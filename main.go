package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFile := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFile)
	check(err, fmt.Sprintf("Failed to open the CSV file: %s", *csvFile))

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	check(err, "Failed to open the provided CSV file")

	problems := ParseLines(lines)

	correct := 0
	for i, p := range problems {

		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d!\n", correct, len(problems))
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}
