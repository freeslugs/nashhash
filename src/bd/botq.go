package bd

import (
	"crypto/ecdsa"
	"log"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
// The locking here has to be careful, otherwise a serious bottleneck can arise.
func (bq *BotQ) Dispatch(number uint, address string) error {

	//// LOCK
	bq.qLock.Lock()
	var wg sync.WaitGroup

	// TODO: we might want to error out here?
	botn := 0
	if int(number) > len(bq.ready) {
		log.Printf("WARNING BotQ.Dispatch %f ether: ready %d, need %d\n", bq.guaranteedBalance, len(bq.ready), number)
		botn = len(bq.ready)
	} else {
		botn = int(number)
	}

	// Dispatch individual bots
	toDispatch, ready := bq.ready[:botn], bq.ready[botn:len(bq.ready)]
	bq.ready = ready

	// We can unlock here because we removed
	bq.qLock.Unlock()
	// UNLOCK
	//
	//

	// Dont have to lock anything here, because these bots are not going to be available
	// to other routines as we removed them from bq.ready
	// Dispatch the bots in parallel
	for i := 0; i < botn; i++ {

		// We send each bot in its goroutine to hopefully speedup the process
		wg.Add(1)
		go func(bot *Bot) {

			e := bot.DoBotStuff(address)
			if e != nil {
				log.Printf("WARNING BotQ.Dispatch %f ether: DoBotStuff failed\n", bq.guaranteedBalance)
			}
			wg.Done()
		}(toDispatch[i])
	}
	wg.Wait()

	// Lets see if any bot is below the guaranteed balance, if so, move it into the refill Q
	var forRefill []*Bot
	var moreReady []*Bot
	for i := 0; i < botn; i++ {

		bot := toDispatch[i]

		bal, err := bot.PendingBalance()
		if err != nil {
			log.Printf("WARNING BotQ.Dispatch %f ether: failed to get bot balance\n", bq.guaranteedBalance)
			forRefill = append(forRefill, bot)
			continue
		}

		// Some bots will be in for refill, some others will not
		if bal < bq.guaranteedBalance {
			forRefill = append(forRefill, bot)
		} else {
			// Back to ready Q
			moreReady = append(moreReady, bot)
		}

	}

	// Now we have to return the bots we borrowed into the Qs,  thus the lock
	// LOCK
	bq.qLock.Lock()
	bq.refill = append(bq.refill, forRefill...)
	bq.ready = append(bq.ready, moreReady...)
	bq.qLock.Unlock()
	// UNLOCK
	//
	//

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

			}
			bq.qLock.Unlock()

			time.Sleep(10 * time.Second)
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
func (bq *BotQ) Kill(harvestAddr common.Address) error {

	bq.qLock.Lock()
	defer bq.qLock.Unlock()

	bq.supervisorDead <- true

	// Kill all the bots
	for _, bot := range bq.ready {
		bot.Kill(harvestAddr)
	}

	for _, bot := range bq.pending {
		bot.Kill(harvestAddr)
	}

	for _, bot := range bq.refill {
		bot.Kill(harvestAddr)
	}

	log.Printf("INFO BotQ %f: dead\n", bq.guaranteedBalance)
	return nil
}
