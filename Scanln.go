package main

import "fmt"

func maim() {
	var name string
	var age int
	n, e := fmt.Scanln(&name, &age)
	fmt.Println("name", name)
	fmt.Println("age", age)
	fmt.Println("numeber of item successfully scanned", n)
	fmt.Println("error", e)
}
