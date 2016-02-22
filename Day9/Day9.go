package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

type routeText string
type path struct{
    to,from string
}
func (r routeText) parse() (routeA, routeB string, distance int){
    words := strings.Split(strings.TrimSpace(string(r))," ")
    d,_ := strconv.Atoi(words[4])
    return words[0],words[2],d
}

func loadFile(folder,file string) string{
	data, err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/AdventOfCode/"+folder+"/"+file)
    if err != nil {
        panic(err)
    }
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
	input := loadFile("Day9","input.txt")
	routes := make(map[path]int)
    cities := []string{}
    for _, str := range strings.Split(input,"\n"){
		route:=routeText(str)
        a,b,d := route.parse()
        routes[path{a,b}] = d
        routes[path{b,a}] = d
        if inSlice(cities,a)==false{
            cities = append(cities,a)
        }
        if inSlice(cities,b)==false{
            cities = append(cities,b)
        }
    }
    permutations := make([][]int, 0)
    items := make([]int,0)
    for i := range cities{
        items = append(items,i)
    }
    getPermutations(items,[]int{},&permutations)
    shortestDistance := 1000000
    longestDistance := 0
    for _,a := range permutations{
        d:=0
        for j:=0;j<len(a)-1;j++{
            from := cities[a[j]]
            to := cities[a[j+1]]
            d += routes[path{to,from}]
        }
        if d<shortestDistance{
            shortestDistance = d
        }
        if d>longestDistance{
            longestDistance = d
        }
    }
    fmt.Println("Part 1", shortestDistance)
    fmt.Println("Part 2", longestDistance)

}


