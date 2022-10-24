package main

import (
    "fmt"
    "math/rand"
    "time"
)
type Game struct {
    // needs to communicate with the game loop
    isRunning bool
    isPaused bool
    // by using a map we can acces the player [*Player] by its name (string = [key])
        /// and then modify the player but this can only happend with a pointer because the pointer
        // can modify it/ Pointer only 8 bytes
    // needs to communicate with the game loop
    players map[string] *Player
    // this quite channel is of type bool meaning that when this is called/used then it must take a bool as an argument
    // needs to communicate with the game loop
    quitCh chan bool
    pauseCh chan bool

}

type Player struct {
    Name string
    AttackPower uint
    Health uint
}
func (g *Game) start(){
    g.isRunning = true
    g.isPaused = false
    g.gameLoop()
}
func(g *Game) addNewPlayer(p *Player){
    // we tell the game (g *Game) give me the players map (g.players (in Game struct))
    // we then store a player in the map using the string to identify the player
    g.players[p.Name] = p
    // prints out the format of newly created player
    fmt.Printf("\nAdding new player %s %d\n ", p.Name, p.Health)
}

func newGame() *Game {
    return &Game{
        isRunning: true,
        isPaused: false,
        // will initlaize map as empty
        players: make(map[string]*Player),
        // channel are like pipes that go to to a particlar place/ used to communicate through Goroutines
        quitCh:  make(chan bool),
    }
    // need to add players
}

func (g *Game) gameLoop(){
    // is equal to 1 second occurance in program execuation
    interval := time.Second * 1
    // is basicall ran in another layer
        // each second (time.Second * 1) it will shoot off this ticker
    ticker := time.NewTicker(interval)
    fmt.Println("the game is running")

    // THE RUNNING IS A NAME FOR THE THE FIRST SWITCH LOOP (ENDS HERE)
    running:
    for {
        select{
            // this basically means, if someone uses this g.quitCh in the game loop
        case <- g.quitCh:
            // it will stop the game loop
                // THE RUNNING IS A NAME FOR THE THE FIRST SWITCH LOOP (ENDS HERE)
            break running
        case <- g.pauseCh:
            g.isPaused = true
            // shoots of ticker
        case <- ticker.C:
            fmt.Println("Paused")
            break
        }
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
    playerB := Player{"Anne", 422,122672}

    game.addNewPlayer(&playerA)
    game.addNewPlayer(&playerB)
    // creates a go function
    // a function that is put in a different layer, that will shoot off when needed

    go game.quitGame()

    game.start()

}
    
func quitGame(quitCh chan bool)  {
    time.Sleep(time.Second * 3)
    quitCh <- true
}
