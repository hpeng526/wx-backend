package server

import (
	"fmt"
	"github.com/hpeng526/wx-backend/mq"
	"github.com/hpeng526/wx/cache"
	"github.com/hpeng526/wx/context"
)

type MqServer struct {
	Ctx context.Context
	Mq  mq.MessageQueue
}

func NewRedisMqServer(appId string, appSecret string, serverAddr string) *MqServer {
	ctx := context.Context{AppID: appId, AppSecret: appSecret, Cache: cache.NewRedisCache(serverAddr)}
	redisMq := mq.NewRedisMq(serverAddr)
	return &MqServer{Ctx: ctx, Mq: redisMq}
}

func (ms *MqServer) HandleMessage(msg string) {
	if msg != "" {
		fmt.Printf("got msg: %s \n", msg)
	}

}
