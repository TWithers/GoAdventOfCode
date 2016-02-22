package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

type Present struct{
	l, w, h, surfaceArea int
}
func (p Present) LW() int{
	return p.l*p.w
}
func (p Present) LH() int{
	return p.l*p.h
}
func (p Present) WH() int{
	return p.w*p.h
}
func (p *Present) SurfaceArea() int{
	p.surfaceArea = 2*p.LW() + 2*p.LH() + 2*p.WH()
	if p.LW() <= p.LH() && p.LW() <= p.WH(){
		p.surfaceArea += p.LW()
	} else if p.LH() <= p.LW() && p.LH() <= p.WH(){
		p.surfaceArea += p.LH()
	} else {
		p.surfaceArea += p.WH()
	}
	return p.surfaceArea
}
func (p Present) RibbonLength() int{
	smallest := min(p.l,p.w)
	if smallest==p.l{
		smallest += min(p.w,p.h)
	}else{
		smallest += min(p.l,p.h)
	}
	return smallest*2
}
func (p Present) BowLength() int{
	return p.l * p.h * p.w
}
func (p Present) TotalRibbon() int{
	return p.RibbonLength() + p.BowLength()
}

func min(x, y int) int{
	if x<y{
		return x
	}else{
		return y
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
	input := loadFile("Day2","input.txt")
	lines := strings.Split(strings.TrimSpace(input),"\n")
	wrappingPaper := 0
	ribbon := 0
	for _,l := range lines{
		l := strings.Split(strings.TrimSpace(l),"x")
		p :=&Present{}
		p.l,_ = strconv.Atoi(l[0])
		p.w,_ = strconv.Atoi(l[1])
		p.h,_ = strconv.Atoi(l[2])
		wrappingPaper+=p.SurfaceArea()
		ribbon+=p.TotalRibbon()
	}
	fmt.Println("Wrapping Paper",wrappingPaper)
	fmt.Println("Ribbon",ribbon)
}