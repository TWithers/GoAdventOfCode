package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)
type sue struct {
	id, children, cats, samoyeds, pomeranians, akitas, vizslas, goldfish, trees, cars, perfumes int
}
func (s sue) part1(real sue) bool{
	if s.children != -1 && s.children != real.children{
		return false
	}
	if s.cats != -1 && s.cats != real.cats{
		return false
	}
	if s.samoyeds != -1 && s.samoyeds != real.samoyeds{
		return false
	}
	if s.pomeranians != -1 && s.pomeranians != real.pomeranians{
		return false
	}
	if s.akitas != -1 && s.akitas != real.akitas{
		return false
	}
	if s.vizslas != -1 && s.vizslas != real.vizslas{
		return false
	}
	if s.goldfish != -1 && s.goldfish != real.goldfish{
		return false
	}
	if s.trees != -1 && s.trees != real.trees{
		return false
	}
	if s.cars != -1 && s.cars != real.cars{
		return false
	}
	if s.perfumes != -1 && s.perfumes != real.perfumes{
		return false
	}
	return true
}
func (s sue) part2(real sue) bool{
	if s.children != -1 && s.children != real.children{
		return false
	}
	if s.cats != -1 && s.cats <= real.cats{
		return false
	}
	if s.samoyeds != -1 && s.samoyeds != real.samoyeds{
		return false
	}
	if s.pomeranians != -1 && s.pomeranians >= real.pomeranians{
		return false
	}
	if s.akitas != -1 && s.akitas != real.akitas{
		return false
	}
	if s.vizslas != -1 && s.vizslas != real.vizslas{
		return false
	}
	if s.goldfish != -1 && s.goldfish >= real.goldfish{
		return false
	}
	if s.trees != -1 && s.trees <= real.trees{
		return false
	}
	if s.cars != -1 && s.cars != real.cars{
		return false
	}
	if s.perfumes != -1 && s.perfumes != real.perfumes{
		return false
	}
	return true
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

func main() {
	input := loadFile("Day16","input.txt")
	lines := strings.Split(input,"\n")
	//var aunts []sue
	realSue := sue{id:0,children: 3,cats: 7,samoyeds: 2,pomeranians: 3,akitas: 0,vizslas: 0,goldfish: 5,trees: 3,cars: 2,perfumes: 1}
	var part1, part2 int
	for i,l := range lines{
		words := strings.Split(l," ")
		s := sue{i+1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1}
		check := []int{2,4,6}
		for _,k := range check{
			val := strings.TrimSpace(words[k+1])
			val = strings.Trim(val,",")
			num,_ := strconv.Atoi(val)
			switch words[k]{
				case "children:":
					s.children = num
				case "cats:":
					s.cats = num
				case "samoyeds:":
					s.samoyeds = num
				case "pomeranians:":
					s.pomeranians = num
				case "akitas:":
					s.akitas = num
				case "vizslas:":
					s.vizslas = num
				case "goldfish:":
					s.goldfish = num
				case "trees:":
					s.trees = num
				case "cars:":
					s.cars = num
				case "perfumes:":
					s.perfumes = num
			}
		}
		if s.part1(realSue){
			part1 = s.id
		}
		if s.part2(realSue){
			part2 = s.id
		}
	}
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}