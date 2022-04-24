package main

import "fmt"

type Gamer struct {
	Name     string
	GameList []string
}

func (g *Gamer) AddGame(name string) {
	g.GameList = append(g.GameList, name)
}

func main() {
	gamer := Gamer{
		Name: "Zelda",
	}

	gamer.AddGame("Mario")
	gamer.AddGame("Fifa 2020")
	gamer.AddGame("Soccer 2020")

	fmt.Println(gamer)
	for _, game := range gamer.GameList {
		fmt.Println(game)
	}
}