package main

import (
	"fmt"
	"math"
	"time"
)
func main() {
	var  num int64
	num = 400000
	begin := time.Now()
	findPrimeBySqrt(num)
	fmt.Println("It took ", time.Since(begin))

	begin = time.Now()
	findPrimeBySix(num)
	fmt.Println("It took ", time.Since(begin))
}

//prime search by
func findPrimeBySqrt(num int64) {
	var i int64
	count := 2
	//fmt.Println("2")
	//fmt.Println("3")
	for i = 5; i <= num; i++ {
		if (isPrime(i)) {
			//fmt.Println(i)
			count++
		}
	}
	fmt.Println("i : ", i-5)
	fmt.Printf("prime count : %d \n", count)
}

//monodidymus prime search
// just detemine if 6n-1 and 6n+1 are prime.
func findPrimeBySix(num int64) {
	var i int64
	count := 2
	//fmt.Println("2")
	//fmt.Println("3")
	maxMultiple := num/6
	for i = 1; i < maxMultiple; i++ {
		if (isPrime(6 * i - 1)) {
			//fmt.Println(6 * i - 1)
			count++
		}
		if (isPrime(6 * i + 1)) {
			//fmt.Println(6 * i + 1)
			count++
		}
	}
	//To avoid boundary detection in every increment
	if (6 * i + 1) <= num {
		if (isPrime(6 * i - 1)) {
			//fmt.Println(6 * i - 1)
			count++
		}
		if (isPrime(6 * i + 1)) {
			//fmt.Println(6 * i + 1)
			count++
		}

	}
	fmt.Println("i : ", i)
	fmt.Printf("prime count : %d \n", count)
}
func isPrimeBySix(num int64) bool {
	var i int64
	tmp := float64(num)
	sqrtVal := int64(math.Sqrt(tmp))
	for i = 5; i <= sqrtVal; i+=6 {
		if (num % i) == 0 {
			return false
		}
	}
	return true
}
func isPrime(num int64) bool {
	var i int64
	tmp := float64(num)
	sqrtVal := int64(math.Sqrt(tmp))
	for i = 2; i <= sqrtVal; i++ {
		if (num % i) == 0 {
			return false
		}
	}
	return true
}
