package gm

// Clerk is client side interface for the GameMaster.
// Still debating if this does more harm than good.
type Clerk struct {
	GMAddr string
}

// Init initializes the Clerk to a GN address
// func (gmk *Clerk) Init(gmAddr string) {
// 	gmk.gmAddr = gmAddr
// }

// ConnectGame connects a game at @gameAddr to the gamemaster.
// By default, GM will start managing the states of the game
// after succesful connection
func (gmk *Clerk) ConnectGame(gameAddr string) error {
	e := connectGame(gmk.GMAddr, gameAddr)
	return e
}

// DisconnectGame disconnect a game at @gameAddr from the master
func (gmk *Clerk) DisconnectGame(gameAddr string) error {
	e := disconnectGame(gmk.GMAddr, gameAddr)
	return e
}
