// Client code
package flyweight

import (
	"fmt"
	"testing"
)

func Test_Flyweight(t *testing.T) {
	game := newGame()

	player1 := game.addTerrorist(TerroristDressType)
	player1.newLocation(1, 1)
	player2 := game.addTerrorist(TerroristDressType)
	player2.newLocation(2, 2)

	player3 := game.addCounterTerrorist(CounterTerroristDressType)
	player3.newLocation(3, 3)
	player4 := game.addCounterTerrorist(CounterTerroristDressType)
	player4.newLocation(4, 4)

	player5 := game.addTerrorist(TerroristDressType)
	player5.newLocation(5, 5)

	t.Log("player1 : " + fmt.Sprintf("%p", player1))
	t.Log("player2 : " + fmt.Sprintf("%p", &player2))
	t.Log("player3 : " + fmt.Sprintf("%p", &player3))
	t.Log("player4 : " + fmt.Sprintf("%p", &player4))
	t.Log("player5 : " + fmt.Sprintf("%p", &player5))

	if fmt.Sprintf("%p", player1.dress) != fmt.Sprintf("%p", player2.dress) {
		t.Error("Not same memory address")
	}

	if fmt.Sprintf("%p", player3.dress) != fmt.Sprintf("%p", player4.dress) {
		t.Error("Not same memory address")
	}
}
