package bd

import (
	"log"
	"sync"
)

// BotQ maintains the bots in three different queues: available, busy and refill
// Remember, bot is an interface.
type BotQ struct {

	// Lock for consistency
	qLock sync.Mutex

	// The bot qs
	ready   []*Bot
	pending []*Bot
	refill  []*Bot

	guaranteedBalance float64
}

// Dispatch sends the bots to the bots address
func (bq *BotQ) Dispatch(number uint, address string) error {
	bq.qLock.Lock()
	defer bq.qLock.Unlock()

	return nil
}

// Init creates nbots, starts the refill goroutine
func (bq *BotQ) Init(guaranteedBalance float64, nbots uint) error {
	bq.qLock.Lock()
	defer bq.qLock.Unlock()

	bq.guaranteedBalance = guaranteedBalance

	// Create all the bots and put them into the refill que
	for i := uint(0); i < nbots; i++ {

		// Create a new bot
		bot := &Bot{}
		e := bot.Init()
		if e != nil {
			log.Fatal(e)
		}

		bq.refill = append(bq.refill, bot)

	}

	return nil
}
