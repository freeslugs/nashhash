package gm

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"net/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// IMPORTANT CONFIGS
const (
	// GameOperator Stuff
	// DisconnectOperator command to stop a game
	DisconnectOperator = 0
	// EthClientPath is the path to the geth.ipc
	EthClientPath = "/Users/me/Library/Ethereum/rinkeby/geth.ipc"

	// BotDispatcher stuff
	NetworkID             = 4
	MinimumBalanceInStake = 5
	RefillAmountInStake   = 20
)

type ExecuteCallArgs struct {
	Message string
}

type ExecuteCallReply struct {
	Response string
}

// ConnectCallArgs is a struct of arguments to the GM.Connect RPC
type ConnectCallArgs struct {
	ContractAddress string
}

// ConnectCallReply is the reply from GM.Connect RPC
type ConnectCallReply struct {
	Reply string
}

// DisconnectCallArgs is a struct of arguments to the GM.Connect RPC
type DisconnectCallArgs struct {
	ContractAddress string
}

// DisconnectCallReply is the reply from GM.Disconnect RPC
type DisconnectCallReply struct {
	Reply string
}

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
func harvestAccounts(keys []*ecdsa.PrivateKey, ownerAddr common.Address) error {

	for _, privk := range keys {

		err := sendEth(privk, ownerAddr, nil)
		if err != nil {
			return err
		}
	}
	return nil
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

	log.Println(nonce)

	// If the value is nil, we want to transfer the whole balance.
	if value == nil {

		value = big.NewInt(0)
		tax := big.NewInt(0)
		tax.Mul(gasPrice, big.NewInt(21000))
		money, err := conn.BalanceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey), nil)
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
