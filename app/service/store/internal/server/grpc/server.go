package grpc

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	v1 "go-common/app/service/store/api"
	"go-common/app/service/store/internal/service"
)

type grpcServer struct {
	svr *service.Service
}

// New new a grpc server.
func New(svc *service.Service) (ws *warden.Server, err error) {
	var (
		cfg warden.ServerConfig
		ct paladin.TOML
	)
	if err = paladin.Get("grpc.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	ws = warden.NewServer(&cfg)
	v1.RegisterStoreSrvServer(ws.Server(), &grpcServer{svr:svc})
	ws, err = ws.Start()
	return
}


