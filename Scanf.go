package main

import "fmt"

func main() {
	var f float64
	var d int
	n, e := fmt.Scanf("%E %d", &f, &d)
	fmt.Println("read floate ", f)
	fmt.Println("read int", d)
	fmt.Println("number of items successfully scanned ", n)
	fmt.Println("error ", e)
}