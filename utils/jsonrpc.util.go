package utils

import (
	"my-assets-be/types"
	"os"

	"github.com/ybbus/jsonrpc/v2"
)

func CreateNewEthRequest(action string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	rpcClient := jsonrpc.NewClient(os.Getenv("INFURA_HOST"))
	return rpcClient.Call(action, params)
}

func GetEthBlock(blockNumber string) (*types.BlockInfo, error) {
	blockInfo := &types.BlockInfo{}

	rpc, err := CreateNewEthRequest("eth_getBlockByNumber", blockNumber, true)

	if err != nil {
		return blockInfo, err
	}

	if err := rpc.GetObject(&blockInfo); err != nil || blockInfo == nil {
		return blockInfo, err
	}

	return blockInfo, nil
}
