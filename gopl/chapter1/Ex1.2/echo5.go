//Exercise 1.2
package main

import (
	"fmt"
	"os"
)

func main(){
	for idx,val := range os.Args[1:]{
		fmt.Printf("argument [%d] = %s\n",idx,val)
	}
}