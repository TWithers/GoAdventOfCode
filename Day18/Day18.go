package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


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
	input := loadFile("Day18","input.txt")
	
	grid := [][]bool{}
	
	rows := strings.Split(input,"\n")
	for _, rowData := range rows {
		row := []bool{}
		rowData = strings.TrimSpace(rowData)
		cols := strings.Split(rowData,"")
		for _,c := range cols {
			if c=="."{
				row = append(row,false)
			}else{
				row = append(row,true)
			}
		}
		grid = append(grid,row)
	}
	for i:=0;i<100;i++ {
		part1Step(&grid)
	}
	count := 0
	for _, row := range grid{
		for _, val := range row{
			if val {
				count++
			}
		}
	}
	fmt.Println("Part 1: ",count)

	part2Grid := [][]bool{}
	rows = strings.Split(input,"\n")
	for _, rowData := range rows {
		row := []bool{}
		rowData = strings.TrimSpace(rowData)
		cols := strings.Split(rowData,"")
		for _,c := range cols {
			if c=="."{
				row = append(row,false)
			}else{
				row = append(row,true)
			}
		}
		part2Grid = append(part2Grid,row)
	}
	part2Grid[0][0]=true
	part2Grid[len(part2Grid)-1][0]=true
	part2Grid[0][len(part2Grid[0])-1]=true
	part2Grid[len(part2Grid)-1][len(part2Grid[0])-1]=true
	for i:=0;i<100;i++ {
		part2Step(&part2Grid)
	}
	count = 0
	for _, row := range part2Grid{
		for _, val := range row{
			if val {
				count++
			}
		}
	}
	fmt.Println("Part 2: ",count)

}

func part1Step(g *[][]bool){
	newGrid := make([][]bool,len(*g))
	for i := range newGrid {
		newGrid[i] = make([]bool, len((*g)[i]))
	}
	for rowIndex, row := range *g{
		for colIndex, v := range row{
			n := neighbors(rowIndex, colIndex, *g)
			if v {
				newGrid[rowIndex][colIndex] = (n==2 || n==3)
			}else {
				newGrid[rowIndex][colIndex] = (n==3)
			}
		}
	}
	*g = newGrid
}

func part2Step(g *[][]bool){
	newGrid := make([][]bool,len(*g))
	for i := range newGrid {
		newGrid[i] = make([]bool, len((*g)[i]))
	}
	for rowIndex, row := range *g{
		for colIndex, v := range row{
			n := neighbors(rowIndex, colIndex, *g)
			if v {
				newGrid[rowIndex][colIndex] = (n==2 || n==3)
			}else {
				newGrid[rowIndex][colIndex] = (n==3)
			}
		}
	}
	newGrid[0][0]=true
	newGrid[len(newGrid)-1][0]=true
	newGrid[0][len(newGrid[0])-1]=true
	newGrid[len(newGrid)-1][len(newGrid[0])-1]=true
	*g = newGrid
}

func neighbors(a int, b int, g [][]bool) int{
	count := 0
	for i:=-1;i<=1;i++{
		if a+i >= 0 && a+i < len(g) {
			for j:=-1;j<=1;j++{
				if b+j >=0 && b+j < len(g[a]) && g[a+i][b+j]{
					if i==0 && j==0 {
						continue
					}else{
						count++
					}
				}
			}
		}
	}
	return count
}