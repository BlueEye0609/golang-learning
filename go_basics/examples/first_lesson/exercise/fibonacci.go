package main

import "fmt"

func main() {
	n := 10
	fmt.Printf("斐波那契数列, %d 项\n", n)
	for i := 0; i < n; i++ {
		fmt.Println(recur_fibo(i))
	}
}

func recur_fibo(n int) int {
	if n <= 1 {
		return n
	} else {
		return recur_fibo(n-1) + recur_fibo(n-2)
	}
}
