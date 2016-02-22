package main

import (
	"fmt"
	"strings"
	"strconv"
)


func next(input string) string{
	split := strings.Split(input,"")
	var last string
	var counter int
	recombine := []string{}
	for _,v := range split{
		if last != v{
			if last!= ""{
				recombine = append(recombine,strconv.Itoa(counter))
				recombine = append(recombine,last)
			}
			last = v
			counter = 1
		}else{
			counter++
		}
	}
	recombine = append(recombine,strconv.Itoa(counter))
	recombine = append(recombine,last)
	return strings.Join(recombine,"")
}
func main() {
	input := "1321131112"
	part1 := 0
	for i:=0;i<50;i++{
		input = next(input)
		if i==39{
			part1 = len(input)
		}
	}
	fmt.Println("Part 1",part1)
	fmt.Println("Part 2",len(input))
	
}


