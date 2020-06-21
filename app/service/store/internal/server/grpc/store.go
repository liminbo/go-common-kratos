package grpc

import(
	"context"
	v1 "go-common/app/service/store/api"
)

func (g *grpcServer) StoreList(ctx context.Context, req *v1.StoreListReq) (rep *v1.StoreListRep, err error) {
	rep,err = g.svr.StoreList(ctx, req)
	return
}

func (g *grpcServer) OfflineStoreDetail(context.Context, *v1.StoreDetailReq) (*v1.OfflineStoreDetailRep, error) {
	panic("implement me")
}

func (g *grpcServer) OnlineStoreDetail(context.Context, *v1.StoreDetailReq) (*v1.OnlineStoreDetailRep, error) {
	panic("implement me")
}

func (g *grpcServer) FuzzySearchStore(context.Context, *v1.FuzzySearchStoreReq) (*v1.FuzzySearchStoreRep, error) {
	panic("implement me")
}

func (g *grpcServer) AddStore(context.Context, *v1.EditStoreReq) (*v1.AddStoreRep, error) {
	panic("implement me")
}

func (g *grpcServer) EditStore(context.Context, *v1.EditStoreReq) (*v1.EditStoreRep, error) {
	panic("implement me")
}

func (g *grpcServer) DeleteStore(context.Context, *v1.DeleteStoreReq) (*v1.DeleteStoreRep, error) {
	panic("implement me")
}

func (g *grpcServer) StoreImageList(context.Context, *v1.StoreDetailReq) (*v1.StoreImageListRep, error) {
	panic("implement me")
}

func (g *grpcServer) StoreVideoList(context.Context, *v1.StoreDetailReq) (*v1.StoreVideoListRep, error) {
	panic("implement me")
}
