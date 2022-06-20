package cluster

import (
	"fmt"
	"net/rpc"
)

/********************************************************
* @author: Ihc
* @date: 2022/6/20 0020 14:52
* @version: 1.0
* @description:
*********************************************************/

type RPCClient struct {
	client *rpc.Client
}

func (rpc *RPCClient) Stop() {
	err := rpc.client.Close()
	if err != nil {
		return
	}
}

func NewRPCClient(host string, port int) *RPCClient {
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil
	}
	return &RPCClient{client: client}
}
