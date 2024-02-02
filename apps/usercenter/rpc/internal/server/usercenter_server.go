// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package server

import (
	"context"

	"zinx-zero/apps/usercenter/rpc/internal/logic"
	"zinx-zero/apps/usercenter/rpc/internal/svc"
	"zinx-zero/apps/usercenter/rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UsercenterServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UsercenterServer) GetUserAuthByAuthKey(ctx context.Context, in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {
	l := logic.NewGetUserAuthByAuthKeyLogic(ctx, s.svcCtx)
	return l.GetUserAuthByAuthKey(in)
}

func (s *UsercenterServer) GetUserAuthByAccountId(ctx context.Context, in *pb.GetUserAuthByAccountIdReq) (*pb.GetUserAuthyAccountIdResp, error) {
	l := logic.NewGetUserAuthByAccountIdLogic(ctx, s.svcCtx)
	return l.GetUserAuthByAccountId(in)
}

func (s *UsercenterServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

func (s *UsercenterServer) CreateRole(ctx context.Context, in *pb.CreateRoleReq) (*pb.CreateRoleResp, error) {
	l := logic.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

func (s *UsercenterServer) GetRoleInfo(ctx context.Context, in *pb.GetRoleInfoReq) (*pb.GetRoleInfoResp, error) {
	l := logic.NewGetRoleInfoLogic(ctx, s.svcCtx)
	return l.GetRoleInfo(in)
}
