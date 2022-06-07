package main

import "fmt"

type SalaryCalcualtor interface {
	Calculate() int
}

type DayTimeJob struct {
	empId      int
	payPerHour int
	totalHour  int
}

type NightTimeJob struct {
	empId      int
	payPerHour int
	totalHour  int
	nightPay   int
}

func (d DayTimeJob) Calculate() int {
	return d.payPerHour * d.totalHour
}

func (n NightTimeJob) Calculate() int {
	return (n.payPerHour + n.nightPay) * n.totalHour
}

func totalPay(s []SalaryCalcualtor) int {
	total := 0
	for _, value := range s {
		total = total + value.Calculate()
	}

	return total
}

func showInterface(s []SalaryCalcualtor) {
	for _, value := range s {
		fmt.Printf("Interface type %T, value is %v\n", value, value)
	}
}

func main() {
	alba01 := DayTimeJob{1, 2000, 5}
	alba02 := DayTimeJob{2, 2200, 8}
	alba03 := NightTimeJob{3, 2000, 5, 500}
	alba04 := NightTimeJob{4, 2000, 8, 1000}

	allMember := []SalaryCalcualtor{alba01, alba02, alba03, alba04}
	total := totalPay(allMember)

	fmt.Println("Total Alba Cost:", total)

	fmt.Println("------------------------")

	showInterface(allMember)
}
