package pokemon

import (
	"fmt"
	"learngo/helpers"
	"strconv"
)

type BasicSkills interface {
	Tackle()
	SandAttack()
}

type Pokemon struct {
	Name      string
	Nickname  string
	Hitpoints int
	Attack    int
	Defense   int
	Speed     int
	BasicSkills
}

func (p Pokemon) Tackle() {
	fmt.Println("Tackle!")
}

func (p Pokemon) SandAttack() {
	fmt.Println("Sand Attack!")
}

func New(name string, nickName string) Pokemon {
	return Pokemon{
		Name:      name,
		Nickname:  nickName,
		Hitpoints: 100,
		Attack:    20,
		Defense:   5,
		Speed:     10,
	}
}

func inputStarterChoice(pokemons []string) string {
	fmt.Println("Which Pokemon would you like?")
	fmt.Printf("1 - %s\n", pokemons[0])
	fmt.Printf("2 - %s\n", pokemons[1])
	fmt.Printf("3 - %s\n", pokemons[2])

	numString := helpers.Scanner()
	numInt, _ := strconv.Atoi(numString)

	if numInt < 1 || numInt > len(pokemons) {
		fmt.Println("Invalid choice")
		return ""
	}

	return pokemons[numInt-1]
}

func ChooseStarter() Pokemon {
	pokemons := []string{"Bulbasaur", "Charmander", "Squirtle"}

	var chosenPokemon string

	for chosenPokemon == "" {
		chosenPokemon = inputStarterChoice(pokemons)
	}

	fmt.Printf("You chose %s!\n", chosenPokemon)

	fmt.Printf("Please choose a name for your %s\n", chosenPokemon)
	chosenName := helpers.Scanner()

	return New(chosenPokemon, chosenName)
}
