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

func  newPLayer(name string, hp uint, ap uint) *Player {
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
	game.start()

	fmt.Println()

   //playerA := newPLayer("Bob", 100, 100)
   //playerB := newPLayer("Alice", 150, 100)

	//playerB.dies()
	//playerA.dies()

//    fmt.Println("The Health of the playerA equals to ", playerA.Health)
}
