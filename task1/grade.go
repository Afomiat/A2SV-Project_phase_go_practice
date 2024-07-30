package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func accept(subjects []string, grades []int) map[string]int {
	dict := make(map[string]int)

	for i := 0; i < len(subjects); i++ {
		subj := subjects[i]
		grade := grades[i]

		dict[subj] = grade
	}

	return dict
}

func Avg(d map[string]int, n int) int {
	Avg := 0

	for _, val := range d {
		Avg += val
	}

	return Avg / n
}

func stuGrade(d map[string]int) map[string]string {
	garde_dict := make(map[string]string)

	for key, val := range d {
		switch {
		case val >= 90:
			garde_dict[key] = "A"
		case val >= 80:
			garde_dict[key] = "B"
		case val >= 70:
			garde_dict[key] = "C"
		default:
			garde_dict[key] = "D"
		}
	}
	return garde_dict
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name please :  ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter number of subjects :  ")
	numSub, _ := reader.ReadString('\n')
	numSub = strings.TrimSpace(numSub)

	numS, err := strconv.Atoi(numSub)
	if err != nil {
		fmt.Println("Invalid value:", numS)
		return
	}

	subjects := make([]string, numS)
	grades := make([]int, numS)

	for i := 0; i < numS; i++ {
		fmt.Printf("Enter subject number %v:  ", i+1)
		subj, _ := reader.ReadString('\n')
		subj = strings.TrimSpace(subj)
		subjects[i] = subj

		fmt.Print("Enter Grade:  ")
		grade, _ := reader.ReadString('\n')
		grade = strings.TrimSpace(grade)

		gra, err := strconv.Atoi(grade)
		if err != nil {
			fmt.Println("Invalid value:", grade)
			return
		}

		grades[i] = gra
	}

	dicti := accept(subjects, grades)
	dict_res := stuGrade(dicti)

	fmt.Println()
	fmt.Println("..................GRADE REPORT..............")
	fmt.Printf("Student's Name:  %v \n", name)
	Average := Avg(dicti, numS)

	for key, val := range dict_res {
		fmt.Printf("%v ____________________ %v\n", key, val)
	}

	fmt.Printf("Average:  %v", Average)
}
