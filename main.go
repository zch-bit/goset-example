package main

import (
	"fmt"

	"goset/goset"
)

func main() {
	set := goset.NewSet()
	set.Add("a")
	set.Add("b")
	set.Add("c")
	set.Add("a")
	fmt.Printf("set: %v\n", set)
	set.Remove("a")
	fmt.Printf("set: %v\n", set)
	fmt.Printf("set contains: %v\n", set.Contains("a"))
	fmt.Printf("set contains: %v\n", set.Contains("c"))
	set.Clear()
	fmt.Printf("set: %v\n", set)
}
