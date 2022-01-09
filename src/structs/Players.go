package structs

import "errors"

type Player struct {
	name       string
	properties []Property
	money      int
	turnOrder  int
	isInJail   bool
	isBankrupt bool
}

func NewPlayer(newName string, newTurnOrder int) Player {
	return Player{
		name:      newName,
		turnOrder: newTurnOrder,
		money:     800,
	}
}

func (p Player) Name() string {
	return p.name
}

func (p Player) Money() int {
	return p.money
}

func (p Player) Properties() []Property {
	return p.properties
}

func (p *Player) BuyProperty(property *Property) error {
	if p.money < property.price {
		return errors.New("failed purcase not enough money")
	}
	p.money -= property.price
	property.isOwned = true
	p.properties = append(p.properties, *property)
	return nil
}
