package main

type player struct {
	hp    int
	maxHP int
	//weapon weapon
	x, y   int
	roomID int
}

func (p *Player) move(x, y int) {
	p.x = x
	p.y = y
}

func (p Player) Draw() rune {
	return '@'
}

func newPlayer(x, y int) Player {
	return Player{
		hp:    15,
		maxHP: 15,
		//weapon: newWeapon("UNARMED", 5, '?'),
		x:      x,
		y:      y,
		roomID: 1,
	}
}
