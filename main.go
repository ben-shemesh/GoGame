package main

import (
	"fmt"
	"time"
	// "time"
)

type Player struct {
	// only positive health
    Health uint
    Name string
    AttackPower uint
	Age uint
}
type Game struct {
	isRunning bool
	isPause bool
	// the key is a string
	// the value of a the map is a pointer to Player
	players map[string]*Player
	quitCh chan bool
	pauseCh chan bool
	playerCh chan *Player
	playerStatsCh chan *Player
}
func (g *Game) addPlayer(p *Player){
	// links to the map
	// stores Player by p.Name
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
		playerCh: make(chan *Player),
	}
}

func fightPlayers(p1 *Player, p2 *Player) {
	// var deadPlayer
	fmt.Printf("%s's health is equal to %d\n",p1.Name, p1.Health)
	fmt.Printf("%s's health is equal to %d\n",p2.Name, p2.Health)
	if p1.Health * 2 < p2.Health {
		p2.Health = p2.Health/ 2
		fmt.Printf("%s has lost %d health.\n" ,p2.Name, p2.Health)
		if p2.Health == 0 {
			p2.dies()
			fmt.Printf("%s has died\n", p2.Name)
			// deadPlayer = p2
		}
	} else  {
			fmt.Printf("%s has lost %d health.\n" ,p1.Name, p1.Health)
			if p1.Health == 0 {
				p1.dies()
				fmt.Printf("%s has died\n", p1.Name)
				// deadPlayer = p1 
			}
		}
	// return deadPlayer
}
func (g *Game) start() {
	g.isRunning = true
	g.gameLoop()
}
func (g *Game) gameLoop() {
    interval1 := 1 * time.Second
	ticker := time.NewTimer(interval1)
	fmt.Println("THE LOOP")

	// causes the select to stop but the for loop contiunes to run
	// named loop "running"
		/// >>> starts "running loop"
	running:
	for {
		select{
		case <- g.quitCh:
			// will make a break the select loop
			break running ////// >>>> ends "running" loop
		case <- g.pauseCh:
			g.isPause = true
		case <- ticker.C:
			fmt.Println("the game is running")
		}
	}
	fmt.Println("Bye Bye")
}

func newPlayer(name string, hp uint, ap uint, ag uint ) *Player {
	return &Player{
		Health: hp,
		Name: name,
		AttackPower: ap,
		Age: ag,
		}
}
func (p *Player) dies()  {
	p.Health = 0
}
func stop(quitCh chan bool)  {
	time.Sleep(time.Second * 5)
	quitCh <- true
}
func  addNewPlayer(playersCh  chan *Player) {
	time.Sleep(time.Second * 5)
	player := newPlayer("JoeBuyDem", 50000, 1, 29)
	playersCh <- player
}
func  displayThingsAfterFight(p *Player)  {
	fmt.Printf("ALice's health is %d\n", p.Health)
}
func main() {
	// move the start in a routine  / not correct usage but shows the point of Go routines
	game := newGame()
	playerA := newPlayer("Bob", 100,100, 54)
	playerB := newPlayer("Alice", 2000,101, 23)
	game.addPlayer(playerA)
	game.addPlayer(playerB)
	// go fightPlayers(playerA, playerB)
	// go displayThingsAfterFight(playerB)
	// go quite(game.quitCh)
	 go addNewPlayer(game.playerCh)
	game.start()
	// test 
}
func quite (quitCh chan bool) {
	time.Sleep(time.Second * 5)
	quitCh <- true
}

