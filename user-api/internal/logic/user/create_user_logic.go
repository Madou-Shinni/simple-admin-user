package user

import (
	"context"

	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"
	"simple-admin-user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.UserRpc.CreateUser(l.ctx,
		&user.UserInfo{
			Id:       req.Id,
			Name:     req.Name,
			Gender:   req.Gender,
			Avatar:   req.Avatar,
			Username: req.Username,
			Mobile:   req.Mobile,
			IdCard:   req.IdCard,
			Openpid:  req.Openpid,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
