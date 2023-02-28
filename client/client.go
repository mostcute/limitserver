package client

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github/mostcute/limitserver/types"
	"log"
)

type KmClient struct {
	server  []string
	xclient client.XClient
}

func NewKmClient(server []string) *KmClient {
	kbpair := []*client.KVPair{}
	for _, s := range server {
		kbpair = append(kbpair, &client.KVPair{
			Key: s,
		})
	}
	d, _ := client.NewMultipleServersDiscovery(kbpair)
	xclient := client.NewXClient("Limit", client.Failover, client.RandomSelect, d, client.DefaultOption)
	xclient.Auth("bearer tGzv3JOkF0XG5Qx2TlKWIA")
	return &KmClient{
		server:  server,
		xclient: xclient,
	}
}

func (k *KmClient) Close() {
	err := k.xclient.Close()
	if err != nil {
		log.Println("err", err)
	}
}
func (k *KmClient) GetToken() error {
	args := types.ArgsGetToken{}
	reply := &types.ReplyGetToken{}
	err := k.xclient.Call(context.Background(), runFuncName(), args, reply)
	if err != nil {
	}
	return err
}
