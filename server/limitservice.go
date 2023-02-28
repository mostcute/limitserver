package server

import (
	"context"
	"fmt"
	"github/mostcute/limitserver/types"
)

// the second parameter is not a pointer
func (t *LimitService) GetToken(ctx context.Context, args types.ArgsGetToken, reply *types.ReplyGetToken) error {
	err := t.limiter.Wait(context.Background())
	reply.Res = "hello"
	if err != nil {
		fmt.Println(err)
	}
	return err
}
