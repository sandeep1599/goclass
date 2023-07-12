package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name   string
	Grades []int
}

func main() {
	students := make(map[string]Student)

	for {
		var name string
		var gradesStr string

		fmt.Print("Enter student name (or 'quit' to exit): ")
		fmt.Scanln(&name)

		if strings.ToLower(name) == "quit" {
			break
		}

		fmt.Print("Enter grades for", name, "(separated by spaces): ")
		fmt.Scanln(&gradesStr)

		grades := parseGrades(gradesStr)

		student := Student{
			Name:   name,
			Grades: grades,
		}

		students[name] = student
	}

	fmt.Println("\nStudent grades:")

	for name, student := range students {
		average := calculateAverage(student.Grades)
		fmt.Printf("%s - Average Grade: %.2f\n", name, average)
	}
}

// ["89",  "98",  "78"]
func parseGrades(gradesStr string) []int {
	grades := make([]int, 0)

	gradesArr := strings.Split(gradesStr, ",")

	for i := 0; i < len(gradesArr); i++ {

		intVal, _ := strconv.Atoi(gradesArr[i])

		grades = append(grades, intVal)
	}

	return grades
}

func calculateAverage(grades []int) float64 {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}
