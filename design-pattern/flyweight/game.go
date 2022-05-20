// Client code
package flyweight

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

func (g *game) addTerrorist(dressType string) *player {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)

	return player
}

func (g *game) addCounterTerrorist(dressType string) *player {
	player := newPlayer("CT", dressType)
	g.counterTerrorists = append(g.counterTerrorists, player)

	return player
}