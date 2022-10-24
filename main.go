package main

import (
    "fmt"
    "math/rand"
    "time"
)
type Game struct {
    isRunning bool
    // by using a map we can acces the player [*Player] by its name (string = [key])
        /// and then modify the player but this can only happend with a pointer because the pointer
        // can modify it/ Pointer only 8 bytes
    players map[string] *Player
}

type Player struct {
    Name string
    AttackPower uint
    Health uint
}
func (g *Game) start(){
    g.isRunning = true
    g.gameLoop()
}
func(g *Game) addNewPlayer(p *Player){
    // we tell the game (g *Game) give me the players map (g.players (in Game struct))
    // we then store a player in the map using the string to identify the player
    g.players[p.Name] = p
    // prints out the format of newly created player
    fmt.Printf("Adding new player %s %d\n ", p.Name, p.Health)
}

func newGame() *Game {
    return &Game{
        isRunning: true,
        // will initlaize map as empty
        players: make(map[string]*Player),
    }
    // need to add players
}

func (g *Game) gameLoop(){
    // is equal to 1 second occurance in program execuation
    interval := time.Second * 1
    for {
        fmt.Println("the game is running")
        time.Sleep(interval)
    }
}
func addNewPLayer(pN string, pA uint , pH uint  ) *Player {
    return &Player{
        Name: pN,
        AttackPower: pA,
        Health: pH,
    }
}
func (p *Player) killPlayer(){
	p.Health = 0
}

func (g *Game) quitGame(){
    
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
playerA := Player{"tony", 122,1222}
game.addNewPlayer(&playerA)
game.start()

// firstPLayer := addNewPlayer("Tony Balongi", 100, 100)
// // killPlayer(firstPLayer)
// // fmt.Println(firstPLayer)
// secondPLayer := addNewPLayer("Gloria Turning", 122, 8348)


// fmt.Printf("The players name is %s the players attack power is %d and the players health is %d\n",
//              firstPLayer.Name, firstPLayer.AttackPower,firstPLayer.Health )
// fmt.Printf("The players name is %s the players attack power is %d and the players health is %d\n",
//              secondPLayer.Name, secondPLayer.AttackPower,secondPLayer.Health )
}