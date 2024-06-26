package client

import (
	"context"
	"log"

	"github.com/mostcute/limitserver/types"
	"github.com/smallnest/rpcx/client"
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
func (k *LimitClient) GetToken() (string, error) {
	args := types.ArgsGetToken{}
	reply := &types.ReplyGetToken{}
	err := k.xclient.Call(context.Background(), runFuncName(), args, reply)
	if err != nil {
	}
	//println("Name = ", reply.Res)
	return reply.Res, err
}

func (k *LimitClient) Usage() (string, error) {
	args := types.ArgsGetToken{}
	reply := &types.ReplyGetToken{}
	err := k.xclient.Call(context.Background(), runFuncName(), args, reply)
	if err != nil {
	}
	//println("Name = ", reply.Res)
	return reply.Res, err
}

func (k *LimitClient) Limit() (string, error) {
	args := types.ArgsGetToken{}
	reply := &types.ReplyGetToken{}
	err := k.xclient.Call(context.Background(), runFuncName(), args, reply)
	if err != nil {
	}
	return reply.Res, err
}
