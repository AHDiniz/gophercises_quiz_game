package main

// Problem is a struct that holds information about a quiz
type Problem struct {
	q string
	a string
}

// ParseLines is a function that gets a slice a creates problems based in it
// The slice is in the following format: [question answer]
func ParseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
