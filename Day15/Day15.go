package main

import (
	"fmt"
)
type ingredient struct {
	capacity, durability, flavor, texture, calories int
}
func main() {
	sprinkles:=ingredient{capacity:2, durability:0, flavor:-2, texture:0, calories:3}
	butterscotch:=ingredient{capacity:0, durability:5, flavor:-3, texture:0, calories:3}
	chocolate:=ingredient{capacity:0, durability:0, flavor:5, texture:-1, calories:8}
	candy:=ingredient{capacity:0, durability:-1, flavor:0, texture:5, calories:8}
	total := 100
	part1 := 0
	part2 := 0
	for a:=0; a<=total; a++{
		remainingA := total-a
		for b:=0; b<=remainingA; b++{
			remainingB := remainingA - b
			for c:=0; c<=remainingB; c++{
				d := remainingB - c
				capacity := (a*sprinkles.capacity)+(b*butterscotch.capacity)+(c*chocolate.capacity)+(d*candy.capacity)
				durability := (a*sprinkles.durability)+(b*butterscotch.durability)+(c*chocolate.durability)+(d*candy.durability)
				flavor := (a*sprinkles.flavor)+(b*butterscotch.flavor)+(c*chocolate.flavor)+(d*candy.flavor)
				texture := (a*sprinkles.texture)+(b*butterscotch.texture)+(c*chocolate.texture)+(d*candy.texture)
				calories := (a*sprinkles.calories)+(b*butterscotch.calories)+(c*chocolate.calories)+(d*candy.calories)
				if capacity<0{
					capacity=0
				}
				if durability<0{
					durability=0
				}
				if flavor<0{
					flavor=0
				}
				if texture<0{
					texture=0
				}
				if calories<0{
					calories=0
				}
				score := capacity*durability*flavor*texture
				if score>part1 {
					part1=score
				}
				if calories==500 && score>part2 {
					part2=score
				}
				
			}
		}
	}
	fmt.Println("Part 1: ",part1)
	fmt.Println("Part 2: ",part2)
}