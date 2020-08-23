package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 直接打印出cpu的逻辑核心数
	fmt.Println("cpu:", runtime.NumCPU())
	
}
