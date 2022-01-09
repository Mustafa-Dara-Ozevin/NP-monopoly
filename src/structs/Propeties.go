package structs

type colors int

type Property struct {
	name       string
	position   int
	price      int
	numHouses  int
	isOwned    bool
	isMortaged bool
	setColor   colors
}

func NewProperty(newName string, newPosition, newPrice int, newSetColor colors) Property {
	return Property{
		name:     newName,
		position: newPosition,
		price:    newPrice,
		setColor: newSetColor,
	}
}

func (p Property) Name() string {
	return p.name
}
