package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

type Santa struct{
	x, y int
}
func (s Santa) visit(h map[string]int) map[string]int{
	location := strconv.Itoa(s.x)+"x"+strconv.Itoa(s.y)
	_, isset := h[location]
	if isset{
		h[location]++
	}else{
		h[location]=1
	}
	return h
}

func loadFile(folder,file string) string{
	data, err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/AdventOfCode/"+folder+"/"+file)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func main() {
	input := loadFile("Day3","input.txt")
	directions := strings.Split(strings.TrimSpace(input),"")
	santa := &Santa{0,0}
	p2Santa := &Santa{0,0}
	p2RoboSanta := &Santa{0,0}
	houses := make(map[string]int)
	p2houses := make(map[string]int)
	houses = santa.visit(houses)
	p2houses = p2Santa.visit(p2houses);
	p2houses = p2RoboSanta.visit(p2houses);
	for i,v := range directions{
		switch v{
		case "^":
			santa.y--
			if i%2==0{
				p2Santa.y--
			}else{
				p2RoboSanta.y--
			}
		case "v":
			santa.y++
			if i%2==0{
				p2Santa.y++
			}else{
				p2RoboSanta.y++
			}
		case "<":
			santa.x--
			if i%2==0{
				p2Santa.x--
			}else{
				p2RoboSanta.x--
			}
		case ">":
			santa.x++
			if i%2==0{
				p2Santa.x++
			}else{
				p2RoboSanta.x++
			}
		}
		houses = santa.visit(houses)
		if i%2==0{
			p2houses = p2Santa.visit(p2houses)
		}else{
			p2houses = p2RoboSanta.visit(p2houses)
		}
	}
	fmt.Println("Part 1 Houses:",len(houses))
	fmt.Println("Part 2 Houses:",len(p2houses))
}