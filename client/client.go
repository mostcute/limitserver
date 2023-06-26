package client

import (
	"context"
	"github.com/mostcute/limitserver/types"
	"github.com/smallnest/rpcx/client"
	"log"
)

type LimitClient struct {
	server  []string
	xclient client.XClient
}

func NewLimitClient(server []string) *LimitClient {
	kbpair := []*client.KVPair{}
	for _, s := range server {
		kbpair = append(kbpair, &client.KVPair{
			Key: s,
		})
	}
	d, _ := client.NewMultipleServersDiscovery(kbpair)
	xclient := client.NewXClient("Limit", client.Failover, client.RandomSelect, d, client.DefaultOption)
	xclient.Auth("bearer tGzv3JOkF0XG5Qx2TlKWIA")
	return &LimitClient{
		server:  server,
		xclient: xclient,
	}
}

func (k *LimitClient) Close() {
	err := k.xclient.Close()
	if err != nil {
		log.Println("err", err)
	}
}
func (k *LimitClient) GetToken() error {
	args := types.ArgsGetToken{}
	reply := &types.ReplyGetToken{}
	err := k.xclient.Call(context.Background(), runFuncName(), args, reply)
	if err != nil {
	}
	println("Name = ", reply.Res)
	return err
}

func (k *LimitClient) Limit() error {
	args := types.ArgsGetToken{}
	reply := &types.ReplyGetToken{}
	err := k.xclient.Call(context.Background(), runFuncName(), args, reply)
	if err != nil {
	}
	println("Limit = ", reply.Res)
	return err
}
