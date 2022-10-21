package main
import "fmt"

type Player struct {
	// only positive health
    Health uint
    Name string
    AttackPower uint
}
type  Game struct {

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
   playerA := newPLayer("Bob", 100, 100)
   playerB := newPLayer("Alice", 150, 100)

	playerA.dies()
	playerB.dies()

    fmt.Println("The Health of the playerA equals to ", playerA.Health)
}
