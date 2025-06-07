package main

import "fmt"

func getGrade(score int) string {
	switch {
	case score >= 80:
		return "A"
	case score >= 60:
		return "B"
	case score >= 40:
		return "C"
	case score >= 20:
		return "D"
	default:
		return "F"
	}
}

func main() {
	students := []struct {
		id    int
		score int
	}{
		{1, 89},
		{2, 65},
		{3, 45},
		{4, 30},
		{5, 10},
	}

	for _, student := range students {
		grade := getGrade(student.score)
		fmt.Printf("Student %d has Grade : %s\n", student.id, grade)
	}
}
