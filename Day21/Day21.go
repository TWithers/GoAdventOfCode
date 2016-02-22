package main

import "fmt"

type person struct{
	life, damage, armor int
}

type item struct{
	name string
	cost, damage, armor int
}

func main() {
	//me := person{life:100,damage:0, armor:0}
	boss := person{life:100,damage:8,armor:2}
	var weapons []item
	weapons = append(weapons,item{"Dagger",8,4,0})
	weapons = append(weapons,item{"Shortsword",10,5,0})
	weapons = append(weapons,item{"Warhammer",25,6,0})
	weapons = append(weapons,item{"Longsword",40,7,0})
	weapons = append(weapons,item{"Greataxe",74,8,0})
	var armors []item
	armors = append(armors,item{"nothing",0,0,0})
	armors = append(armors,item{"Leather",13,0,1})
	armors = append(armors,item{"Chainmail",31,0,2})
	armors = append(armors,item{"Splintmail",53,0,3})
	armors = append(armors,item{"Bandedmail",75,0,4})
	armors = append(armors,item{"Platemail",102,0,5})
	var rings []item
	rings = append(rings,item{"NoRing 1",0,0,0})
	rings = append(rings,item{"NoRing 2",0,0,0})
	rings = append(rings,item{"Damage +1",25,1,0})
	rings = append(rings,item{"Damage +1",25,1,0})
	rings = append(rings,item{"Damage +2",50,2,0})
	rings = append(rings,item{"Damage +3",100,3,0})
	rings = append(rings,item{"Defense +1",20,0,1})
	rings = append(rings,item{"Defense +2",40,0,2})
	rings = append(rings,item{"Defense +3",80,0,3})
	var combos [][]item
	minCost := 1000
	maxCost := 0
	var itemSet1 []item 
	var itemSet2 []item 
	combos = createCombos(weapons,armors,rings)
	for _,c := range combos{
		attack := c[0].damage+c[2].damage+c[3].damage - boss.armor;
		if attack < 1 {
			attack = 1
		}
		defense := boss.damage - (c[1].armor+c[2].armor+c[3].armor);
		if defense < 1{
			defense = 1
		}
		cost := c[0].cost+c[1].cost+c[2].cost+c[3].cost
		if attack >=defense{
			if cost < minCost{
				minCost = cost
				itemSet1 = c
			}
		}else{
			if cost > maxCost{
				maxCost = cost
				itemSet2 = c
			}
		}
	}
	fmt.Println("Part 1: ",minCost,itemSet1)
	fmt.Println("Part 2: ",maxCost,itemSet2)
}

func createCombos (weapons []item,armors []item, rings []item) [][]item{
	var c [][]item
	for _,w := range weapons{
		for _,a := range armors{
			for i1,r1 := range rings{
				for i2,r2 := range rings{
					if i1==i2{
						continue
					}
					equipment := []item{w,a,r1,r2}
					c = append(c, equipment)
				}
			}
		}
	}
	return c
}