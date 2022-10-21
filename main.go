package main

import (
    "fmt"
    "time"
)

type Player struct {
	// only positive health
    Health uint
    Name string
    AttackPower uint
}
type Game struct {
	isRunning bool
	// the key is a string
	// the value of a the map is a pointer to Player
	players map[string]*Player
}
func (g *Game) addPlayer(p *Player){
	// links to the map
	// stores Player by its name
	g.players[p.Name] = p
	fmt.Println("adding new players")

}
func newGame() *Game {
    return &Game{
		isRunning: false,
	}
}
func (g *Game) start() {
	g.isRunning = true
	g.gameLoop()
}
func (g *Game) gameLoop() {
    interval1 := 1 * time.Second
	for  {
		fmt.Println("the game is running")
		//
		time.Sleep(interval1)
    }
}

func  newPlayer(name string, hp uint, ap uint ) *Player {
	return &Player{
		Health: hp,
		Name: name,
		AttackPower: ap,
		}
}
func (p *Player) dies()  {
	p.Health = 0
}

func main() {
	game := newGame()
	// move the start in a routine  / not correct usage but shows the point of Go routines
	go game.start()
	playerA := newPlayer("Bob", 100,100)
	game.addPlayer(playerA)


	fmt.Printf("%+v", game)

}
