package main

import (
	"fmt"
	"sort"
)

func main() {
	containers := []int{33,14,18,20,45,35,16,35,1,13,18,13,50,44,48,6,24,41,30,42}
	sort.Sort(sort.Reverse(sort.IntSlice(containers)))
	p1 := part1(0,containers,0)
	p2 := []int{}
	part2(0,containers,0,&p2)
	sort.Ints(p2)
	count:=0
	min := 0
	for _,v := range p2{
		if min==0{
			min=v
		}
		if v==min{
			count++
		}else{
			break
		}

	}

	fmt.Println("Part 1: ",p1)
	fmt.Println("Part 2: ",count)


}

func part1(sum int, containers []int, count int) int{
	for i,v := range containers {
		tempSum := sum
		tempSum += v
		if tempSum==150 {
			count++
		}else if tempSum < 150 && i != len(containers)-1 { 
			count += part1(tempSum,containers[i+1:],0)
		}
	}
	return count
}
func part2(sum int, containers []int, steps int,results *[]int){
	steps++
	for i,v := range containers {
		tempSum := sum
		tempSum += v
		if tempSum==150 {
			*results = append(*results,steps)
		}else if tempSum < 150 && i != len(containers)-1 { 
			part2(tempSum,containers[i+1:],steps, results)
		}
	}
}