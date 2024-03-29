package server

import (
	"context"
	"fmt"
	"github.com/mostcute/limitserver/types"
	"time"
)

// the second parameter is not a pointer
func (t *LimitService) GetToken(ctx context.Context, args types.ArgsGetToken, reply *types.ReplyGetToken) error {
	//err := t.limiter.Wait(context.Background())
	for i := 0; i < len(t.limiters); i++ {
		if i == (len(t.limiters) - 1) {
			//fmt.Println("wait ", t.limiters[i].Name)
			if t.limiters[i].Limiter.Allow() {
				reply.Res = t.limiters[i].Name
				return nil
			} else {
				i = -1
				time.Sleep(time.Millisecond * 10)
				continue
			}
		} else {
			if t.limiters[i].Limiter.Allow() {
				reply.Res = t.limiters[i].Name
				return nil
			} else {
				continue
			}
		}
	}
	return nil
}

// the second parameter is not a pointer
func (t *LimitService) Limit(ctx context.Context, args types.ArgsGetToken, reply *types.ReplyGetToken) error {
	//err := t.limiter.Wait(context.Background())
	total := 0.0
	for i := 0; i < len(t.limiters); i++ {
		total += t.limiters[i].Limit
	}
	reply.Res = fmt.Sprintf("%f", total)
	return nil
}
