package main

import (
    "fmt"
    "math/rand"
    "time"
)
type Game struct {
    isRunning bool
}
func (g *Game) start(){
    g.gameLoop()
    g.isRunning = true
}

type Player struct {
    Name string
    AttackPower uint
    Health uint
}
func addNewPLayer(pN string, pA uint , pH uint  ) *Player {
    return &Player{
        Name: pN,
        AttackPower: pA,
        Health: pH,
    }
}
func killPlayer(p *Player){
	p.Health = 0
}

func newGame() *Game {
    return &Game{
        isRunning: true,
    }
}
func (g *Game) gameLoop(){
    // is equal to 1 second occurance in program execuation
    interval := time.Second * 1
    timer := time.NewTimer(interval)
    for {
		fmt.Println("looping")
		<- timer.C
		timer.Reset(interval)
	}
	
}
func (p *Player) powerDepleter() {
    rand.Seed(time.Now().Unix())
    rando := rand.Intn(5)
    p.AttackPower = p.AttackPower / uint(rando)
}
func (p *Player) powerBooster(){
    rand.Seed(time.Now().Unix())
    rando := rand.Intn(5)
    p.Health = p.Health * uint(rando)
}

func main (){
game := newGame()
game.start()
firstPLayer := addNewPLayer("Tony Balongi", 100, 100)
// killPlayer(firstPLayer)
// fmt.Println(firstPLayer)
secondPLayer := addNewPLayer("Gloria Turning", 122, 8348)


fmt.Printf("The players name is %s the players attack power is %d and the players health is %d\n",
             firstPLayer.Name, firstPLayer.AttackPower,firstPLayer.Health )
fmt.Printf("The players name is %s the players attack power is %d and the players health is %d\n",
             secondPLayer.Name, secondPLayer.AttackPower,secondPLayer.Health )
}