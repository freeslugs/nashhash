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

	multiplier := 3.0

	// We shall lock it before anything else
	bq.qLock.Lock()
	// We only need to do something if the refill q is not empty
	if len(bq.refill) != 0 {

		var refill []*Bot
		for i := 0; i < len(bq.refill); i++ {

			//log.Println("in refill loop")

			bot := bq.refill[i] //bq.refill[len(bq.refill)-1], bq.refill[:len(bq.refill)-1]

			// Send money
			e := sendEth(refillKey, bot.auth.From, toWei(bq.guaranteedBalance*multiplier))
			if e != nil {
				refill = append(refill, bot)
				log.Printf("WARNING BotQ.Refill %f ether: %s\n", bq.guaranteedBalance, e)

			} else {
				// Push onto the pending q
				bq.pending = append(bq.pending, bot)
				log.Printf("INFO BotQ.Refill %f ether: refill succesful 0x%x", bq.guaranteedBalance, bot.auth.From)
			}

		}

		bq.refill = refill

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

				// Every Bot's current actual balance has to be checked
				var pending []*Bot

				for _, bot := range bq.pending {

					bal, err := bot.Balance()
					if err != nil {
						log.Println(err)
						continue
					}

					if bal >= bq.guaranteedBalance {
						bq.ready = append(bq.ready, bot)
					} else {
						pending = append(pending, bot)
					}
				}

				// Restore the bots who's payment is still pending
				bq.pending = pending

				log.Printf("INFO BotQ.supervisor %f ether: %d bots pending\n",
					bq.guaranteedBalance,
					len(bq.pending))

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
	bq.supervisorDead = make(chan bool)
	go bq.supervisor()

	return nil
}

// Kill BotQ
func (bq *BotQ) Kill() error {

	bq.supervisorDead <- true

	return nil
}
