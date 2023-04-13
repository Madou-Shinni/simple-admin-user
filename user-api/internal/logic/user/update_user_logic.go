package user

import (
	"context"

	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"
	"simple-admin-user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.UserRpc.UpdateUser(l.ctx,
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
