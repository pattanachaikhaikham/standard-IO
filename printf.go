package main

import "fmt"

func main() {
	fmt.Printf("MY name is %s, I am %d years old\n", "ice", 19)

	fmt.Printf("type = %T \n", 3.14159265359)
	fmt.Printf("pi=%f\n", 3.14159265359)
	fmt.Printf("pi=%2.2f\n", 3.14159265359)
	fmt.Printf("pi=%9.f\n", 3.14159265359)
	fmt.Printf("pi=%-9.f\n", 3.14159265359)
	fmt.Printf("pi=%09.f\n", 3.14159265359)
	fmt.Printf("pi=%09.2f\n", 3.14159265359)
	fmt.Printf("pi=%E\n", 3.14159265359)

	fmt.Printf("1 > 2 = %t \n", 1 > 2)
	fmt.Printf("10 (base 2) = %b \n", 10)
	fmt.Printf("10 (base 2) = %o \n", 10)
	fmt.Printf("10 (base 2) = %d \n", 10)
	fmt.Printf("10 (base 2) = %x \n", 10)
}
