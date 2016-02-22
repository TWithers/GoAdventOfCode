package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)


type direction struct{
	text string
}
func (d direction) parse(cache map[string]uint16, wires map[string]wire) (map[string]uint16, map[string]wire){
	words := strings.Split(d.text," ")
	if len(words)==3{
		v,notInt := strconv.ParseUint(words[0],10,16)
		if notInt!=nil{
			wires[words[2]] = wire{b:words[0]}
		}else{
			cache[words[2]] = uint16(v)
		}
	}else if len(words)==4{
		w := wire{operator:"NOT"}
		v,notInt := strconv.ParseUint(words[1],10,16)
		if notInt!=nil{
			w.b =words[1]
		}else{
			w.bInt=uint16(v)
		}
		wires[words[3]]=w
	}else{
		w := wire{operator:words[1]}
		v,notInt := strconv.ParseUint(words[0],10,16)
		if notInt!=nil{
			w.a =words[0]
		}else{
			w.aInt=uint16(v)
		}
		vv,notInta := strconv.ParseUint(words[2],10,16)
		if notInta!=nil{
			w.b =words[2]
		}else{
			w.bInt=uint16(vv)
		}
		wires[words[4]]=w
	}
	return cache, wires
}
type wire struct{
	operator,a,b string
	aInt,bInt uint16
}
var tabs int
func findWire(name string, cache *map[string]uint16, wires map[string]wire) uint16{
	v,isset := (*cache)[name]
	if isset != false{
		return v
	}
	w:=wires[name]
	if w.operator!="NOT"{
		if len(w.a)!=0{
			w.aInt = findWire(w.a,cache, wires)
		}
	}
	if len(w.b)!=0{
		w.bInt = findWire(w.b,cache,wires)
	}
	var ret uint16
	switch w.operator{
	case "NOT":
		ret = ^w.bInt
	case "AND":
		ret = w.aInt & w.bInt
	case "OR":
		ret = w.aInt | w.bInt
	case "LSHIFT":
		ret = w.aInt << uint16(w.bInt)
	case "RSHIFT":
		ret = w.aInt >> uint16(w.bInt)
	default:
		ret = w.bInt
	}
	(*cache)[name]=ret
	return ret
}
func loadFile(folder,file string) string{
	data, err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/AdventOfCode/"+folder+"/"+file)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func main() {
	input := loadFile("Day7","input.txt")
	directions := strings.Split(strings.TrimSpace(input),"\n")
	cache := make(map[string]uint16)
	part2cache := make(map[string]uint16)
	wires := make(map[string]wire)
	for _,v := range directions{
		d := direction{text:strings.TrimSpace(v)}
		cache, wires = d.parse(cache,wires)
		part2cache, wires = d.parse(part2cache,wires)
	}
	a := findWire("a",&cache,wires)
	fmt.Println("Part 1",a)
	part2cache["b"] = a
	part2a := findWire("a",&part2cache,wires)
	fmt.Println("Part 2", part2a)
}