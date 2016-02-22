package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)
type reindeer struct {
	name string
	speed, travelTime, restTime, timer, distanceTraveled, points int
	isTravelling bool
}
func (r *reindeer) advance(){
	if r.timer == 0 {
		r.isTravelling = !r.isTravelling
		if r.isTravelling {
			r.timer = r.travelTime
		}else{
			r.timer = r.restTime
		}
	}
	r.timer--
	if r.isTravelling{
		r.distanceTraveled += r.speed
	}
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
	input := loadFile("Day14","input.txt")
	lines := strings.Split(input,"\n")
	var deer []reindeer
	for _,l := range lines{
		words := strings.Split(l," ")
		speed,_ := strconv.Atoi(words[3])
		travelTime,_ := strconv.Atoi(words[6])
		restTime,_ := strconv.Atoi(words[13])
		deer = append(deer,reindeer{name:words[0],speed:speed,travelTime:travelTime,restTime:restTime,timer:0,distanceTraveled:0,points:0,isTravelling:false})
	}
	var part1, part2 int
	for i:=0;i<2503;i++{
		for i := range deer {
			deer[i].advance()
		}
		part1 = 0
		for _,d := range deer{
			if d.distanceTraveled>part1{
				part1 = d.distanceTraveled
			}
		}
		for i := range deer {
			if deer[i].distanceTraveled == part1 {
				deer[i].points++
			}
		}
	}
	part2 = 0
	for _,d := range deer{
		if d.points>part2{
			part2 = d.points
		}
	}
	
	fmt.Println("Part 1: ",part1)
	fmt.Println("Part 2: ",part2)
}