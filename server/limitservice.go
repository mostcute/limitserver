package server

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/mostcute/limitserver/types"
)

// the second parameter is not a pointer
func (t *LimitService) GetToken(ctx context.Context, args types.ArgsGetToken, reply *types.ReplyGetToken) error {

	defer func() {
		atomic.AddUint64(&t.Used, 1)
	}()
	//err := t.limiter.Wait(context.Background())
	for i := 0; i < len(t.limiters); i++ {
		if i == (len(t.limiters) - 1) {
			//fmt.Println("wait ", t.limiters[i].Name)
			if t.limiters[i].Limiter.Allow() {
				reply.Res = t.limiters[i].Name
				return nil
			} else {
				i = -1
				time.Sleep(time.Millisecond * 1)
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

// the second parameter is not a pointer
func (t *LimitService) Usage(ctx context.Context, args types.ArgsGetToken, reply *types.ReplyGetToken) error {
	//err := t.limiter.Wait(context.Background())
	// total := atomic.LoadUint64(&t.Used5min)
	count := t.M.Count()
	reply.Res = fmt.Sprintf("last call %d \n count %d  \n 1-min rate: %f \n 5-min rate: %f \n 15-min rate: %f  \n mean rate: %f \n", count-t.Last, count, t.M.Rate1(), t.M.Rate5(), t.M.Rate15(), t.M.RateMean())
	t.Last = count
	return nil
}
