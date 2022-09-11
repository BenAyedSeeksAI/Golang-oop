package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func instanceCalculate(list [][]string) []Calculate {
	var ElementCalculate Calculate
	var listCalculate []Calculate
	for _, line := range list {
		age, _ := strconv.Atoi(line[2])
		Salary, _ := strconv.Atoi(line[3])
		Bonus, _ := strconv.Atoi(line[4])
		if line[0] == "Engineer" {
			ElementCalculate = Engineer{
				details: Employee{Name: line[1], Age: age, Salary: Salary, Bonus: Bonus},
			}
		} else if line[0] == "Technician" {
			ElementCalculate = Technician{
				details: Employee{Name: line[1], Age: age, Salary: Salary, Bonus: Bonus},
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
		Element := strings.Split(fileScanner.Text(), " ")
		list = append(list, Element)
	}
	readFile.Close()
	return list
}

func main() {
	var fileName string = "data.txt"
	list := readFileData(fileName)
	listOfSalaries := instanceCalculate(list)
	for _, element := range listOfSalaries {
		fmt.Println(element.getProfession())
		fmt.Println(element.annualSalary())
	}
}
