package grpc

import(
	"context"
	v1 "go-common/app/service/store/api"
)

func (g grpcServer) StoreList(context.Context, *v1.StoreListReq) (*v1.StoreListRep, error) {
	panic("implement me")
}

func (g grpcServer) StoreDetail(context.Context, *v1.StoreDetailReq) (*v1.StoreDetailRep, error) {
	panic("implement me")
}

func (g grpcServer) FuzzySearchStore(context.Context, *v1.FuzzySearchStoreReq) (*v1.FuzzySearchStoreRep, error) {
	panic("implement me")
}

func (g grpcServer) AddStore(context.Context, *v1.AddStoreReq) (*v1.AddStoreRep, error) {
	panic("implement me")
}

func (g grpcServer) EditStore(context.Context, *v1.EditStoreReq) (*v1.EditStoreRep, error) {
	panic("implement me")
}

func (g grpcServer) StoreAutoReply(context.Context, *v1.StoreAutoReplyReq) (*v1.StoreAutoReplyRep, error) {
	panic("implement me")
}

func (g grpcServer) QueryStoreListByContactUser(context.Context, *v1.QueryStoreListByContactUserReq) (*v1.QueryStoreListByContactUserRep, error) {
	panic("implement me")
}

func (g grpcServer) DeleteStore(context.Context, *v1.DeleteStoreReq) (*v1.DeleteStoreRep, error) {
	panic("implement me")
}
