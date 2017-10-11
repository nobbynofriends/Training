package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("decimal: %d \t  binary: %b \t hexadecimal: %x \t UTF-8: %q \n", i, i, i, i)
	}
}
