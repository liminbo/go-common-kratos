package grpc

import(
	"context"
	v1 "go-common/app/service/store/api"
)

func (g *grpcServer) StoreList(ctx context.Context, req *v1.StoreListReq) (rep *v1.StoreListRep, err error) {
	rep,err = g.svr.StoreList(ctx, req)
	return
}

func (g *grpcServer) OfflineStoreDetail(ctx context.Context, req *v1.StoreDetailReq) (rep *v1.OfflineStoreDetailRep, err error) {
	rep,err = g.svr.OfflineStoreDetail(ctx, int(req.StoreId))
	return
}

func (g *grpcServer) OnlineStoreDetail(ctx context.Context, req *v1.StoreDetailReq) (rep *v1.OnlineStoreDetailRep, err error) {
	rep,err = g.svr.OnlineStoreDetail(ctx, int(req.StoreId))
	return
}

func (g *grpcServer) FuzzySearchStore(ctx context.Context, req *v1.FuzzySearchStoreReq) (rep *v1.FuzzySearchStoreRep, err error) {
	rep,err = g.svr.StoreFuzzySearch(ctx, req)
	return
}

func (g *grpcServer) AddStore(ctx context.Context, req *v1.EditStoreReq) (rep *v1.AddStoreRep, err error) {
	storeId,err := g.svr.AddStore(ctx, req)

	rep = new(v1.AddStoreRep)
	rep.StoreId = int32(storeId)
	return
}

func (g *grpcServer) EditStore(ctx context.Context, req *v1.EditStoreReq) (rep *v1.EditStoreRep, err error) {
	rep = new(v1.EditStoreRep)
	err = g.svr.EditStore(ctx, req)
	return
}

func (g *grpcServer) DeleteStore(ctx context.Context, req *v1.DeleteStoreReq) (rep *v1.DeleteStoreRep, err error) {
	rep = new(v1.DeleteStoreRep)
	err = g.svr.DeleteStore(ctx, int(req.StoreId))
	return
}

func (g *grpcServer) StoreImageList(ctx context.Context, req *v1.StoreDetailReq) (rep *v1.StoreImageListRep, err error) {
	rep,err = g.svr.StoreImageList(ctx, int(req.StoreId))
	return
}

func (g *grpcServer) StoreVideoList(ctx context.Context, req *v1.StoreDetailReq) (rep *v1.StoreVideoListRep, err error) {
	rep,err = g.svr.StoreVideoList(ctx, int(req.StoreId))
	return
}
