package main

/*
#cgo LDFLAGS: ./lib/mylib.o -lstdc++
#include "./lib/mylib.h"
*/
import "C"
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
    a := 5
    b := 3
    result := C.add(C.int(a), C.int(b))
    fmt.Printf("Resultado da adição: %d\n", int(result))
	resultSub := C.sub(C.int(a), C.int(b))
	fmt.Printf("Resultado da subtração: %d\n", int(resultSub))
}

func init() {
	//check if g++ is installed
	if _, err := exec.LookPath("g++"); err != nil {
		fmt.Println("g++ is not installed")
		fmt.Println("Please install g++ and try again")
		fmt.Println("On windows, you can install MinGW with: choco install mingw")
		os.Exit(1)
	}
	// build the c++ library
	// g++ -c lib/mylib.cpp -o lib/mylib.o
	cmd := exec.Command("g++", "-c", "lib/mylib.cpp", "-o", "lib/mylib.o")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error building the c++ library")
		os.Exit(1)
	}
}
