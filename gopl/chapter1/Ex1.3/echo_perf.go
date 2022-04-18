//Exercise 1.3
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end := time.Now()
	fmt.Printf("Time taken using strings.Join = %v\n", end.Sub(start))
	start1 := time.Now()
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end1 := time.Now()
	fmt.Printf("Time taken using For range = %v\n", end1.Sub(start1))

}
