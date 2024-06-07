package client

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc/metadata"
)

func NewGRPCConnection(c zrpc.RpcClientConf) zrpc.Client {
	client := zrpc.MustNewClient(c)
	return client
}

func WithAuthToken(ctx context.Context, tkn string) context.Context {
	md := metadata.New(map[string]string{"authorization": "Bearer " + tkn})
	return metadata.NewOutgoingContext(ctx, md)
}
