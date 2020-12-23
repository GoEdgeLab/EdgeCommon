package dao

type BaseDAO struct {
}

func (this *BaseDAO) RPC() RPCClient {
	return sharedRPCClient
}
