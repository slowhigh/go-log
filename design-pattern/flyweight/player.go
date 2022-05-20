// Context
package flyweight

type player struct {
	dress dress
	playerType string
	latitude int
	longitude int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress: dress,
	}
}

func (p *player) newLocation(lat, long int) {
	p.latitude = lat
	p.longitude = long
}