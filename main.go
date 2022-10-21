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
	isPause bool
	// the key is a string
	// the value of a the map is a pointer to Player
	players map[string]*Player
	quitCh chan bool
	pauseCh chan bool
	PlayerCh chan *Player
}
func (g *Game) addPlayer(p *Player){
	// links to the map
	// stores Player by its name
	g.players[p.Name] = p
	fmt.Printf("adding new players %s - %d\n", p.Name, p.AttackPower)
}
func newGame() *Game {
    return &Game{
		isRunning: false,
		isPause: false,
		players: make(map[string]*Player),
		quitCh:  make(chan bool),
		pauseCh: make(chan bool),
		PlayerCh: make(chan *Player),
	}
}
func (g *Game) start() {
	g.isRunning = true
	g.gameLoop()
}

func (g *Game) gameLoop() {
    interval1 := 1 * time.Second
	ticker := time.NewTicker(interval1)
	fmt.Println("the game is running")
	running:
	for  {
        select {
        case player := <- g.PlayerCh:
			g.addPlayer(player)
        case <- g.quitCh:
			break running
        case <- g.pauseCh:
           g.isPause = true
        case <- ticker.C:
		 	fmt.Println("Game is running")

        }
    }
	fmt.Println("Bye Bye")
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
	// move the start in a routine  / not correct usage but shows the point of Go routines
	game := newGame()
	playerA := newPlayer("Bob", 100,100)
	playerB := newPlayer("Alice", 10001,101)
	game.addPlayer(playerA)
	game.addPlayer(playerB)
	go addNewPlayer (game.PlayerCh)
	game.start()

}
func stop(quitCh chan bool)  {
	time.Sleep(time.Second * 5)
	quitCh <- true
}
func  addNewPlayer (playersCh  chan *Player) {
	time.Sleep(time.Second * 5)
	player := newPlayer("JoeBuyDem", 50000, 1)
	playersCh <- player
}