package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"strconv"
	// "reflect"
)
func getSum(input string) int{
	current := ""
	sum := 0
	for i := 0; i<len(input); i++{
		letter := string(input[i])
		if letter=="-" || letter=="1" || letter=="2" || letter=="3" || letter=="4" || letter=="5" || letter=="6" || letter=="7" || letter=="8" || letter=="9" || letter=="0" {
			current += letter
		}else{
			if len(letter)>0 {
				num,_ := strconv.Atoi(current)
				sum += num
				current = ""
			}
		}
	}
	return sum
}
func part2(u interface{}) int {
	sum := 0
	if m,valid := u.(map[string]interface{}); valid {
		for _,v := range m{
			if num,valid := v.(float64); valid {
				sum += int(num)
			}else if str,valid := v.(string);valid{
				if str == "red"{
					return 0
				}
			}else {
				sum += part2(v)
			}
		}
	}else if a,valid := u.([]interface{}); valid{
		for _,v := range a{
			if num,valid := v.(float64);valid{
				sum += int(num)
			}else {
				sum += part2(v)
			}
		}
	}
	return sum
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
	input := loadFile("Day12","input.txt")
	fmt.Println("Part 1: ",getSum(input))
	var j interface{}
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		fmt.Println("Error",err)
	}
	sum := part2(j)
	fmt.Println("Part 2: ",sum)
}