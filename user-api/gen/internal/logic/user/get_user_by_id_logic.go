package user

import (
	"context"

	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"
	"simple-admin-user-rpc/types/user"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.IDReq) (resp *types.UserInfoResp, err error) {
	data, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UserInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:     data.Name,
			Gender:   data.Gender,
			Avatar:   data.Avatar,
			Username: data.Username,
			Mobile:   data.Mobile,
			IdCard:   data.IdCard,
			Openpid:  data.Openpid,
		},
	}, nil
}
