package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var arr = [...]int{1, 2, 3}
	fmt.Println(arr)
	a := [2]int{1, 2}
	a2 := [2]int{2, 1}
	b := [2]int{1, 2}
	fmt.Println(a == a2, a == b)

	blockSeg()

	byteArr := [32]byte{}
	for i, _ := range byteArr {
		byteArr[i] = byte(i + 1)
	}
	fmt.Println(byteArr)
	zeroAnother(byteArr)
	fmt.Println(byteArr)
	zero(&byteArr)
	fmt.Println(byteArr)

	blockSeg()

	months := []string{1: "C", 2: "Java", 3: "Golang"}
	fmt.Printf("%T, %d, %d\n", months, len(months), cap(months))
	months = append(months, "hello")
	fmt.Printf("%T, %d, %d\n", months, len(months), cap(months))

	blockSeg()

	goMap := map[string]int{
		"Hello": 1,
		"World": 2,
	}
	fmt.Println(goMap)
	if hello, ok := goMap["Hello"]; ok {
		fmt.Println(hello)
	}
	delete(goMap, "Hello")
	fmt.Println(goMap)
	// put
	goMap["Test"] = 12
	fmt.Println(goMap)

	blockSeg()

	// 结构体也可以内部定义貌似
	var dilbert Employee
	dilbert.Salary = 5000
	var pEmployee = &dilbert
	pEmployee.ID = 100
	var otherEmp = EmployeeByID(10)
	fmt.Println(otherEmp.ID)
	var anotherEmp = EmployeeByIDOther(11)
	fmt.Println(anotherEmp.ID)
	anotherEmp.ID = 12
	fmt.Println(anotherEmp.ID)
	// err op
	// EmployeeByIDOther(111).ID = 12

	blockSeg()

	// anonymous
	var w Wheel
	w.X = 8 // eq w.Point.X = 8
	w.Point.Y = 10

}

func EmployeeByID(id int) *Employee {
	return &Employee{
		ID: id,
	}
}

func EmployeeByIDOther(id int) Employee {
	return Employee{
		ID: id,
	}
}

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}

func zeroAnother(arr [32]byte) {
	arr = [32]byte{}
}

func blockSeg() {
	fmt.Printf("\n---------------------------\n")
}
