package user

import (
	"context"

	"simple-admin-user-rpc/gen/internal/svc"
	"simple-admin-user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.IDsReq) (*user.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &user.BaseResp{}, nil
}
