package bd

import (
	"sync"
)

// BotDispatcher is a rpc server that allows the client to seamlessly dispatch bots
type BotDispatcher struct {

	// For adding new stakes
	bdLock sync.Mutex

	// A map from stakes to BotQueues
	queues map[float32]*BotQueues
}

// BotQueues maintains the bots in three different queues: available, busy and refill
type BotQueues struct {

	// Lock for consistency
	qLock sync.Mutex

	// The bot ques
	available []*Bot
	busy      []*Bot
	refill    []*Bot
}

type Bot interface {
	DoBotStuff()
}
