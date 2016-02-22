package main

import (
	"fmt"
	"math"
)

type elf struct{
	id, next int
}

func main() {
	// var elves []elf
	// totalPresents := 0
	// for i:=1; totalPresents < 36000000; i++ {
	// // for i:=1; i<10; i++ {
	// 	elves = append(elves, elf{i,i})
	// 	totalPresents = 0
	// 	for elfIndex:=1;elfIndex<len(elves/2);elfIndex++{
	// 		if elves[elfIndex].next == i {
	// 			totalPresents += e.id*10
	// 			elves[elfIndex].next = e.next + e.id
	// 		} 
	// 	}
	// 	// fmt.Println(totalPresents)
	// }
	// fmt.Println(len(elves))


	//You can brute force this, but it will take a VERY long time.  Several minutes or more to run the operation.
	//Best bet is to figure out all the divisors of the number
	//If a number is 16, you only need to count up to the square root (4)
	// Loop 1, 1 and 16 are both divisors
	// Loop 2, 2 and 8 are both divisors
	// Loop 3, nothing
	// Loop 4, 4 is a divisor (do not count the other 4 because it would be a duplicate)
	//at 64, you are only making 8 loops.  at 144=12, 10000=100, this is going to be the most efficient way to calculate divisors
	
	totalPresents := 0
	house := 1
	for i:=1;totalPresents < 36000000;i++{
		sq := math.Sqrt(float64(i))
		totalPresents = 0
		for a:=1;float64(a)<=sq;a++{
			if i%a==0{
				totalPresents += a*10
				if a!=i/a{
					totalPresents += (i/a)*10
				}
			}
		}
		house = i
	}
	fmt.Println("Part 1: ",house)

	totalPresents = 0
	house = 1
	for i:=1;totalPresents < 36000000;i++{
		sq := math.Sqrt(float64(i))
		totalPresents = 0
		for a:=1;float64(a)<=sq;a++{
			if i%a==0{
				if i<a*50{
					totalPresents += a*11
				}
				if a!=i/a && i<(i/a)*50{
					totalPresents += (i/a)*11
				}
			}
		}
		house = i
	}
	fmt.Println("Part 2: ", house)
}