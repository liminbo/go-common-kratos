package v1

import (
	"context"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"google.golang.org/grpc"
)

// AppID .
const AppID = "TODO: ADD APP ID"

// NewClient new grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (StoreSrvClient, error) {
	client := warden.NewClient(cfg, opts...)
	//cc, err := client.Dial(context.Background(), fmt.Sprintf("discovery://default/%s", AppID))
	cc, err := client.Dial(context.Background(), "direct://default/127.0.0.1:9000")
	if err != nil {
		return nil, err
	}
	return NewStoreSrvClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm api.proto
