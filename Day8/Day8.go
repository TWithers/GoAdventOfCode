package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func unquote(str string) string {
    s, _ := strconv.Unquote(str)
    return s
}

func quote(str string) string {
    return strconv.Quote(str)
}

func loadFile(folder,file string) string{
	data, err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/AdventOfCode/"+folder+"/"+file)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func main() {
	input := loadFile("Day8","input.txt")
	// lines := strings.Split(strings.TrimSpace(input),"\n")
	var total int
	for _, str := range strings.Split(input,"\n"){
		//fmt.Println()
        total += len(strings.TrimSpace(str)) - len(unquote(strings.TrimSpace(str)))
    }
    fmt.Println(total)

    // part 2
    total = 0
    for _, str := range strings.Split(input, "\n") {
        total += len(quote(strings.TrimSpace(str))) - len(strings.TrimSpace(str))
    }
    fmt.Println(total)
}