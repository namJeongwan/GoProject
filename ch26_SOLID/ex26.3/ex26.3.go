package main

import "fmt"

type Attacker interface {
	GetName() string
}

type DamageTaker interface {
	DealDamage(attakcer Attacker, damage int)
}

type Player struct {
	name string
}

func (p *Player) GetName() string {
	return p.name
}

type Monster struct {
	hp int
}

func (p *Player) Attack(dt DamageTaker, damage int) {
	dt.DealDamage(p, damage)
}

func (m *Monster) DealDamage(attacker Attacker, damage int) {
	m.hp -= damage
	if m.hp < 0 {
		fmt.Printf("Monster killed by %v\n", attacker.GetName())
	}
}

func main() {
	player := &Player{"TestUser"}
	monster := &Monster{150}

	monster.DealDamage(player, 250)
}
