package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

type Latop struct {
	Computer
}

func NewLatop() IProduct {
	return &Latop{
		Computer: Computer{
			name:  "Laptop",
			stock: 100,
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop",
			stock: 85,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "laptop":
		return NewLatop(), nil
	case "desktop":
		return NewDesktop(), nil
	default:
		return nil, fmt.Errorf("Invalid computer type")
	}
}

func printNameAndStock(computer IProduct) {
	fmt.Printf("Name: %s, Stock: %d\n", computer.getName(), computer.getStock())
}

func main() {
	latop, _ := GetComputerFactory("laptop")
	printNameAndStock(latop)

	desktop, _ := GetComputerFactory("desktop")
	printNameAndStock(desktop)

	_, error := GetComputerFactory("miniPC")
	if error != nil {
		fmt.Println(error)
	}
}
