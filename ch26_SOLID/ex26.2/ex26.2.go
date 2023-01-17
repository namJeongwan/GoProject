package main

import "fmt"

type Attacker interface {
	DealDamage()
}

type Player struct {
	name string
}

type Monster struct {
	hp int
}

func (p *Player) DealDamage(m *Monster, damage int) {
	m.hp -= damage
	if m.hp < 0 {
		fmt.Printf("Monster killed by %v\n", p.name)
	}
}

func main() {
	player := &Player{"TestUser"}
	monster := &Monster{150}

	player.DealDamage(monster, 200)
}
