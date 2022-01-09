package main

import (
	"strconv"

	"www.github.com/Mustafa-Dara-Ozevin/NP-monopoly/src/structs"
)

func main() {
	levent := structs.NewProperty("levent", 1, 300, structs.Navy)
	player1 := structs.NewPlayer("Mustapha", 1)
	println(player1.Name() + " has " + strconv.Itoa(player1.Money()) + "$")
	player1.BuyProperty(&levent)
	println(player1.Name() + " bought " + player1.Properties()[0].Name())
	println(player1.Name() + " has " + strconv.Itoa(player1.Money()) + "$")
}
