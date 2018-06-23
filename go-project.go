package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("C:\\Users\\Lenovo\\Projects\\go\\goproject\\goproject\\mycd.bat")
	thiserr := cmd.Run()
	if thiserr != nil {
		log.Fatal(thiserr)
		fmt.Println(thiserr)
	}

}
