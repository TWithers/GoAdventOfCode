package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
	"math"
)

type lightingGrid [1000][1000]int

func (g *lightingGrid) toggle(d direction){
	for x:=d.startX;x<=d.endX;x++{
		for y:=d.startY;y<=d.endY;y++{
			if d.flag==-1{
				g[x][y] = int(math.Abs(float64(g[x][y]-1)))
			}else{
				g[x][y] = d.flag
			}
		}
	}
}
func (g *lightingGrid) adjust(d direction){
	for x:=d.startX;x<=d.endX;x++{
		for y:=d.startY;y<=d.endY;y++{
			if d.flag==0{
				g[x][y]--
				if g[x][y]<0{
					g[x][y]=0
				}
			}else if d.flag==1{
				g[x][y]++
			}else{
				g[x][y]+=2
			}

		}
	}
}
func (g lightingGrid) count() int{
	count := 0
	for _,x := range g{
		for _,y := range x{
			count += y
		}
	}
	return count
}

type direction struct{
	text string
	flag,startX,startY,endX,endY int
}
func (d *direction) parse(){
	words := strings.Split(d.text," ")
	var start,end string
	if len(words)==4{
		d.flag=-1
		start = words[1]
		end = words[3]
	}else if words[1]=="on"{
		d.flag=1
		start = words[2]
		end = words[4]
	}else{
		d.flag=0
		start = words[2]
		end = words[4]
	}
	s := strings.Split(start,",")
	d.startX,_ = strconv.Atoi(s[0])
	d.startY,_ = strconv.Atoi(s[1])
	e := strings.Split(end,",")
	d.endX,_ = strconv.Atoi(e[0])
	d.endY,_ = strconv.Atoi(e[1])
}


func loadFile(folder,file string) string{
	data, err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/AdventOfCode/"+folder+"/"+file)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func main() {
	input := loadFile("Day6","input.txt")
	directions := strings.Split(strings.TrimSpace(input),"\n")
	var t [1000][1000]int
	part1 := lightingGrid(t)
	part2 := lightingGrid(t)
	for _,v := range directions{
		d := direction{text:strings.TrimSpace(v)}
		d.parse()
		part1.toggle(d)
		part2.adjust(d)
	}
	fmt.Println("Part 1",part1.count())
	fmt.Println("Part 2",part2.count())
}