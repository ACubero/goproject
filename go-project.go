package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Este mensaje nos alerta de que hemos entrado en el programa")
	cmd := exec.Command("C:\\Users\\Lenovo\\Projects\\go\\goproject\\goproject\\mycd.bat")
	thiserr := cmd.Run()
	if thiserr != nil {
		//log.Fatal(thiserr) -- hemos comentado esta línea porque no hacemos ningún log
		fmt.Println(thiserr)
	}
	//Comentario
	fmt.Println("Final")
}
