package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func instanceCalculate(list [][]string) []Calculate {
	var listCalculate []Calculate
	for _, line := range list {
		age, _ := strconv.Atoi(line[2])
		Salary, _ := strconv.Atoi(line[3])
		Bonus, _ := strconv.Atoi(line[4])
		var ElementCalculate Calculate
		if line[0] == "Engineer" {
			ElementCalculate = Engineer{
				details: Employee{Name: line[1],
					Age:    age,
					Salary: Salary,
					Bonus:  Bonus},
			}
		} else if line[0] == "Technician" {
			ElementCalculate = Technician{
				details: Employee{Name: line[1],
					Age:    age,
					Salary: Salary,
					Bonus:  Bonus},
			}
		}
		listCalculate = append(listCalculate, ElementCalculate)

	}
	return listCalculate
}
func readFileData(fileName string) [][]string {
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	list := make([][]string, 0)
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			Element := strings.Split(fileScanner.Text(), " ")
			list = append(list, Element)
		}
	}
	readFile.Close()
	return list
}
func show() {
	var fileName string = "data.txt"
	list := readFileData(fileName)
	listOfSalaries := instanceCalculate(list)
	for _, element := range listOfSalaries {
		fmt.Println(element.getProfession())
		fmt.Println(element.annualSalary())
	}
}
func getArgs() string {
	return os.Args[1]
}
func add() {
	fmt.Println("Enter a whole line of data :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("%T", input)
	// open file
	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	// write the input
	defer file.Close()
	file.WriteString(input + "\n")
}

func main() {
	Argument := getArgs()
	if Argument == "add" {
		add()
	} else if Argument == "show" {
		show()
	}
}
