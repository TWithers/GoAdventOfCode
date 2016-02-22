package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)
type seating struct {
	person, neighbor string
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func loadFile(folder,file string) string{
	data, err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/AdventOfCode/"+folder+"/"+file)
    check(err)
    return string(data)
}
func inSlice(haystack []string, needle string) bool{
    for _,v := range haystack{
        if v==needle{
            return true
        }
    }
    return false
}
func getPermutations(items []int, perm []int, permutations *[][]int) {
    if len(items) == 0 {
        (*permutations) = append((*permutations), perm)
    } else {
        for i := 0; i < len(items); i++ {
            newItems := make([]int, len(items) - 1)
            copy(newItems, items[:i])
            copy(newItems[i:], items[i + 1:])
            newPerm := make([]int, len(perm))
            copy(newPerm, perm)
            getPermutations(newItems, append(newPerm, items[i]),permutations)
        }
    }
}
func main() {
	input := loadFile("Day13","input.txt")
	lines := strings.Split(input,"\n")
	m := make(map[seating]int)
	var people []string
	for _,l := range lines{
		words := strings.Split(l," ")
		t := strings.Trim(strings.TrimSpace(words[10]),".")
		units,_ := strconv.Atoi(words[3])
		if words[2] != "gain" {
			units = 0 - units
		}
		m[seating{words[0],t}] = units
		if !inSlice(people,words[0]){
			people = append(people,words[0])
		}

	}
	// for k,i := range m{
	//  	fmt.Println(k.person,k.neighbor)
	//  	fmt.Println(i)
	// }
	permutations := make([][]int, 0)
	items := make([]int,0)
    for i := range people{
        items = append(items,i)
    }
    getPermutations(items,[]int{},&permutations)
	sum := part1(people, permutations, m)
	fmt.Println("Part 1: ",sum)

	for _,v := range people{
		m[seating{"self",v}]=0
		m[seating{v,"self"}]=0
	}
	people = append(people,"self")
	permutations = make([][]int, 0)
	items = make([]int,0)
    for i := range people{
        items = append(items,i)
    }
    getPermutations(items,[]int{},&permutations)
	sum = part1(people, permutations, m)
	fmt.Println("Part 2: ",sum)
}

func part1(guests []string, p [][]int, m map[seating]int) int{
	highest := 0
	for _,a := range p{
		sum := 0
		for i:=0;i<len(a);i++{
			person := guests[a[i]]
			var next, prev int
			if i == len(a)-1 {
				next = 0
			}else{
				next = i+1
			}
			if i == 0 {
				prev = len(a)-1
			}else{
				prev = i-1
			}

			neighbor := guests[a[next]]
			sum += m[seating{person,neighbor}]
			neighbor = guests[a[prev]]
			sum += m[seating{person,neighbor}]
		}
		// fmt.Println(sum)
		if sum>highest {
			highest = sum
		}
	}
	return highest
}