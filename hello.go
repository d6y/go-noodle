package main

import "fmt"


func seven() int {
	return 7
}

func incr(a func() int) int {
	return a() + 1
}


func main() {
	fmt.Println(incr(seven) )
}

