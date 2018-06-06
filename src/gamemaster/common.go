package gm

type ExecuteCallArgs struct {
	Message string
}

type ExecuteCallReply struct {
	Response string
}

// ConnectCallArgs is a struct of arguments to the GameMaster.Connect RPC
type ConnectCallArgs struct {
	ContractAddress string
}

// ConnectCallReply is the reply from GameMaster.Connect RPC
type ConnectCallReply struct {
	Reply string
}

// DisconnectCallArgs is a struct of arguments to the GameMaster.Connect RPC
type DisconnectCallArgs struct {
	ContractAddress string
}

// DisconnectCallReply is the reply from GameMaster.Disconnect RPC
type DisconnectCallReply struct {
	Reply string
}
