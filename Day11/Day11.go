package main

import (
	"fmt"
	"strings"
)
func inc(s string) string {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] < 122 {
            s = s[:i] + string(s[i]+1) + s[i+1:]
            break
        }
        s = s[:i] + string(97) + s[i+1:]
    }
    b:=[]rune{}
    for i:=0;i<len(s);i++{
    	b=append(b,97)
    }
    if s == string(b){
    	s = "a"+s
    }
    return s
}

type password string

func (s password) hasRun() bool{
	var last rune
	counter := 0
	for _,r := range s{
		last++
		if r == last{
			counter++
		}else{
			counter = 0
			last = r
		}
		if counter==2{
			return true
		}

	}
	return false
}
func (s password) hasPairs() bool{
	var last rune
	pairs := 0
	for _,r := range s{
		if r == last{
			pairs++
			last = 0
		}else{
			last = r
		}
		if pairs==2{
			return true
		}

	}
	return false
}

func (s password) noBadLetters() bool{
	return strings.Index(string(s),"i")==-1 && strings.Index(string(s),"l")==-1 && strings.Index(string(s),"o")==-1
}
func (s password) part1Password() bool{
	return s.hasRun() && s.hasPairs() && s.noBadLetters()
}
func main() {
	input :="hepxcrrq"
	answers :=[]string{}
	for len(answers) <2{
		input = inc(input)
		test := password(input)
		if test.part1Password(){
			answers = append(answers,input)
		}
	}
	fmt.Println("Part 1", answers[0])
	fmt.Println("Part 2", answers[1])
	
	
	
}