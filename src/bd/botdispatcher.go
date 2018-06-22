package bd

import (
	"sync"
)

type BotDispatcher struct {

	// For adding new stakes
	bdLock sync.Mutex

	// A map from stakes to BotQueues
	map[float32]*BotQueues



}

type BotQueues struct {
	
	// Lock for consistency
	qLock sync.Mutex

	// The bot ques 
	available []*Bot
	busy []*Bot
	refill []*Bot

}

type Bot interface {
	DoBotStuff()
}