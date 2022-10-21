package main
import "fmt"

type Player struct {
	// only positive health
    Health uint
    Name string
    AttackPower int
}
type  Game struct {

}
func (p *Player) dies()  {
	p.Health = 0
}

func main() {
    playerA := &Player {
        Health: 100,
        Name: "Bob",
        AttackPower: 50,
    }
	playerB := &Player {
		Health: 150,
		Name: "Bob",
		AttackPower: 50,
		}

	playerA.dies()
	playerB.dies()

    fmt.Println("The Health of the playerA equals to ", playerA.Health)
}
