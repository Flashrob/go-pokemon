package main

import (
	"fmt"
	"learngo/helpers"
	"learngo/player"
	"learngo/pokemon"
)

func main() {

	var isGameRunning bool = true

	for isGameRunning == true {
		fmt.Println("Hello, please enter your name!")
		name := helpers.Scanner()

		player := player.New(name)
		fmt.Println(player)

		starterPokemon := pokemon.ChooseStarter()

		starterPokemon.Tackle()

		isGameRunning = false
	}

}
