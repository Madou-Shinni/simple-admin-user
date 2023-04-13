package user

import (
	"context"

	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"
	"simple-admin-user-rpc/types/user"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (resp *types.UserListResp, err error) {
	data, err := l.svcCtx.UserRpc.GetUserList(l.ctx,
		&user.UserListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
			Gender:   req.Gender,
			Avatar:   req.Avatar,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.UserListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.UserInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Name:     v.Name,
				Gender:   v.Gender,
				Avatar:   v.Avatar,
				Username: v.Username,
				Mobile:   v.Mobile,
				IdCard:   v.IdCard,
				Openpid:  v.Openpid,
			})
	}
	return resp, nil
}
