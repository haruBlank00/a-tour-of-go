package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!!")
	fmt.Println("The time has begun...")
	fmt.Println("TIK - TOC - TIK - TOC", time.Now().Format("15:04:05"))
}
