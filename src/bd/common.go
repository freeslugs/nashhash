package bd

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"net/rpc"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// BotQRefillSleepTime seconds between attempts for refill
	BotQRefillSleepTime = 10

	NetworkID = 4

	// EthClientPath is the path to the geth.ipc
	EthClientPath = "/Users/me/Library/Ethereum/rinkeby/geth.ipc"
)

// DispatchArgs is the argument to the BotDispatcher.Dispatch rpc call
type DispatchArgs struct {
	// ContractAddress is the address of the contract the bot operates on
	ContractAddress string

	// Bot allowance is the balance that the bot has to have to be able
	// to perfoma DoBotStuff
	RequiredBalance float64

	// Number is the number of bots that needs to be dispatched
	Number uint
}

// DispatchReply has the reply
type DispatchReply struct{}

func call(c *rpc.Client, rpcname string,
	args interface{}, reply interface{}) error {

	err := c.Call(rpcname, args, reply)
	if err == nil {
		return nil
	}
	log.Println(err)
	return err
}

// Helper function that harvests balances
func harvestAccount(key *ecdsa.PrivateKey, ownerAddr common.Address) error {

	err := sendEth(key, ownerAddr, nil)
	if err == nil {
		return err
	}

	// If we fail for some reason,
	for attempt := 0; attempt < 5; attempt++ {
		log.Println("WARNING harvestAccount: retrying the harvest")
		err := sendEth(key, ownerAddr, nil)
		if err == nil {
			break
		}
	}

	return nil
}

func sendEthSafe(key *ecdsa.PrivateKey, toAddr common.Address, value *big.Int, lock *sync.Mutex) error {

	lock.Lock()
	defer lock.Unlock()
	return sendEth(key, toAddr, value)

}

func sendEth(key *ecdsa.PrivateKey, toAddr common.Address, value *big.Int) error {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(EthClientPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	// We need to ask the client about currect gas price
	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	// We need to find out the nonce associated with the address
	nonce, err := conn.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		return err
	}

	// If the value is nil, we want to transfer the whole balance.
	if value == nil {

		value = big.NewInt(0)
		tax := big.NewInt(0)
		tax.Mul(gasPrice, big.NewInt(21000))
		money, err := conn.PendingBalanceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey))
		if err != nil {
			return err
		}
		value.Sub(money, tax)
	}

	// This is the transaction to move money
	tx := types.NewTransaction(
		nonce,
		toAddr,
		value,
		21000,
		gasPrice,
		nil)

	// We sign the transaction with the sugarBotKey
	signtx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(NetworkID)), key)
	if err != nil {
		return err
	}

	// Send the transaction into the client
	err = conn.SendTransaction(context.Background(), signtx)
	if err != nil {
		return err
	}
	return nil
}
