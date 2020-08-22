package main

import (
	"fmt"
	
	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/slice"
)

func main() {
	a := slice.IntSlice(make([]int, 0))
	a = append(a, 2)
	a = append(a, 1)
	a = append(a, 3)
	fmt.Printf("%v\n", a)
	
	// sort in ascending
	sort.Sort(a.Begin(), a.End())
	fmt.Printf("%v\n", a)
}
