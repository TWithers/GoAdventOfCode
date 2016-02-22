package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func getHash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	input := "bgvyzdsv"
	found:=false
	part1:=0
	part2:=0
	i:=0
	for found==false{
		test := input + strconv.Itoa(i)
		hash := getHash(test)
		if hash[:5]=="00000" && part1==0{
			part1 = i
		}
		if hash[:6]=="000000" && part2==0{
			part2 = i
			found=true
		}
		i++
	}
	fmt.Println("Part 1",part1)
	fmt.Println("Part 2",part2)
}