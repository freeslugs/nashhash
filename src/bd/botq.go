package bd

import (
	"crypto/ecdsa"
	"log"
	"sync"
)

// BotQ maintains the bots in three different queues: available, busy and refill
// Remember, bot is an interface.
type BotQ struct {

	// Lock for consistency
	qLock sync.Mutex

	// The bot qs
	ready          []*Bot
	pending        []*Bot
	refill         []*Bot
	supervisorDead chan bool

	guaranteedBalance float64
}

// Dispatch sends the bots to the bots address
func (bq *BotQ) Dispatch(number uint, address string) error {
	bq.qLock.Lock()
	defer bq.qLock.Unlock()

	return nil
}

// Refill the BotQ wiht funds. This will move funds to
func (bq *BotQ) Refill(refillKey *ecdsa.PrivateKey) error {

	// We shall lock it before anything else
	bq.qLock.Lock()
	// We only need to do something if the refill q is not empty
	if len(bq.refill) != 0 {

		for i := 0; i < len(bq.refill); i++ {

			// Pop
			var bot *Bot
			bot, bq.refill = bq.refill[len(bq.refill)-1], bq.refill[:len(bq.refill)-1]

			// Send money
			e := sendEth(refillKey, bot.auth.From, toWei(bq.guaranteedBalance*3))
			if e != nil {
				log.Printf("WARNING BotQ.Refill %f ether: %s\n", bq.guaranteedBalance, e)
			}

			// Push onto the pending q
			bq.pending = append(bq.pending, bot)

		}

	}
	bq.qLock.Unlock()
	return nil
}

func (bq *BotQ) supervisor() {

	log.Printf("INFO BotQ %f ether: queue supervisor started\n", bq.guaranteedBalance)

	for {
		select {
		case <-bq.supervisorDead:
			log.Printf("INFO BotQ %f ether: queue supervisor quitting\n", bq.guaranteedBalance)
			return
		default:

			bq.qLock.Lock()
			if len(bq.pending) != 0 {

				for i := 0; i < len(bq.pending); i++ {
					// TODO
				}

			}
			bq.qLock.Unlock()
		}
	}

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

	// Go roputine that manages the Qs
	go bq.supervisor()

	return nil
}
