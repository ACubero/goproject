package main

import (
	"fmt"
	"os"
)

func main() {
	os.Chdir("C:/Users/Lenovo/Projects")

	mydir, err := os.Getwd()
	if err == nil {
		fmt.Println(mydir)
	}
}
