package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type niceString string

func (s niceString) hasDouble() bool{
	letters := strings.Split(strings.TrimSpace(string(s)),"")
	lastLetter :=""
	validDouble :=false
	for _,l := range letters{
		if lastLetter==	l{
			validDouble = true
			lastLetter=""
		}else{
			lastLetter=l
		}
	}
	return validDouble
}

func (s niceString) hasVowels() bool{
	letters := strings.Split(strings.TrimSpace(string(s)),"")
	vowels := 0
	for _,l := range letters{
		if l=="a" || l=="e" || l=="i" || l=="o" || l=="u"{
			vowels++
		}
	}
	return vowels>=3
}

func (s niceString) noBadCharacters() bool{
	if strings.Index(string(s),"ab")==-1 && strings.Index(string(s),"cd")==-1 && strings.Index(string(s),"pq")==-1 && strings.Index(string(s),"xy")==-1{
		return true
	}
	return false
}

func (s niceString) isValidPart1() bool{
	if s.hasDouble() && s.hasVowels() && s.noBadCharacters(){
		return true
	}
	return false
}

func (s niceString) hasPairs() bool{
	letters := strings.Split(strings.TrimSpace(string(s)),"")
	for i := range letters{
		if i==len(letters)-1{
			return false
		}
		index := strings.Index(string(s),letters[i]+letters[i+1])
		lastIndex := strings.LastIndex(string(s),letters[i]+letters[i+1])
		if lastIndex-index>1{
			return true
		}
	}
	return false
}
func (s niceString) hasRepeat() bool{
	letters := strings.Split(strings.TrimSpace(string(s)),"")
	for i,l := range letters{
		if i==len(letters)-2{
			return false
		}
		if l==letters[i+2]{
			return true
		}
    }
	return false
}

func (s niceString) isValidPart2() bool{
	return s.hasPairs() && s.hasRepeat()
}



func loadFile(folder,file string) string{
	data, err := ioutil.ReadFile(os.Getenv("GOPATH")+"/src/AdventOfCode/"+folder+"/"+file)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func main() {
	input := loadFile("Day5","input.txt")
	words := strings.Split(strings.TrimSpace(input),"\n")
	niceCount1 :=0
	niceCount2 :=0
	for _,word := range words{
		s := niceString(word)
		if s.isValidPart1(){
			niceCount1++
		}
		if s.isValidPart2(){
			niceCount2++
		}
	}
	fmt.Println("Part 1",niceCount1)
	fmt.Println("Part 2",niceCount2)
}