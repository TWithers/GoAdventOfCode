package main

import "fmt"

type player struct{
	life, mana, damage, armor int
	effects map[string]effect
}

type effect struct {
	turnsleft int
	task func(*player, *player)
}


func main(){
	boss := player{100,0,8,2,map[string]effect{}}
	mainCharacter := player{100,500,0,0,map[string]effect{}}

	fmt.Println(mainCharacter)
	fmt.Println(boss)

}


func castMagicMissle(p *player, b *boss){
	fmt.Println("Casts Magic Missle for 4 damage!")
	p.mana -= 53
	b.life -= 4
}
func castDrain(p *player, b *boss){
	fmt.Println("Casts Drain and does 2 damage and regenerates 2 life!")
	p.mana -= 73
	p.life += 2
	b.life -= 2
}
func castShield(p *player, b *boss){
	fmt.Println("Casts Shield (increased armor for next ")
	p.mana -= 113
	if _,exists := p.effects["shield"]; !exists{
		p.effects["shield"]=effect{6,shieldEffect}
	}
}
func shieldEffect(p *player, b *boss){
	p.armor = 7
}
func castPoison(p *player, b *boss){
	p.mana -= 173
	if _,exists := b.effects["poison"]; !exists{
		p.effects["poison"]=effect{6,poisonEffect}
	}
}
func poisonEffect(p *player, b *boss){
	b.life -= 3
}
func castRecharge(p *player, b *boss){
	p.mana -= 229
	if _,exists := p.effects["recharge"]; !exists{
		p.effects["recharge"]=effect{5,rechargeEffect}
	}
}
func rechargeEffect(p *player, b *boss){
	p.mana += 101
}


// func main() {
// 	var spells []castable;
// 	spells = append(spells,castable{"Magic Missle",53,4,0,0,0,0})
// 	spells = append(spells,castable{"Drain",73,2,0,2,0,0})
// 	spells = append(spells,castable{"Shield",113,0,7,0,0,6})
// 	spells = append(spells,castable{"Poison",173,3,0,0,0,6})
// 	spells = append(spells,castable{"Recharge",229,0,0,0,101,5})
// 	fmt.Println("Part 1: ",minCost,itemSet1)
// 	fmt.Println("Part 2: ",maxCost,itemSet2)
// }

// func createCombos (weapons []item,armors []item, rings []item) [][]item{
// 	var c [][]item
// 	for _,w := range weapons{
// 		for _,a := range armors{
// 			for i1,r1 := range rings{
// 				for i2,r2 := range rings{
// 					if i1==i2{
// 						continue
// 					}
// 					equipment := []item{w,a,r1,r2}
// 					c = append(c, equipment)
// 				}
// 			}
// 		}
// 	}
// 	return c
// }