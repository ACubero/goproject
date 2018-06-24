package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("C:\\Users\\Lenovo\\Projects\\go\\goproject\\goproject\\mycd.bat")
	thiserr := cmd.Run()
	if thiserr != nil {
		//log.Fatal(thiserr)
		fmt.Println(thiserr)
	}
	//Comentario
	fmt.Println("Final")
}
