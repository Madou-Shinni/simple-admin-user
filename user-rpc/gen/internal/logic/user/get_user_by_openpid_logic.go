package user

import (
	"context"

	"simple-admin-user-rpc/gen/internal/svc"
	"simple-admin-user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByOpenpidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByOpenpidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByOpenpidLogic {
	return &GetUserByOpenpidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByOpenpidLogic) GetUserByOpenpid(in *user.OpenpidReq) (*user.UserInfo, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfo{}, nil
}
