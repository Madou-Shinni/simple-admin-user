package user

import (
	"context"

	"simple-admin-user-rpc/internal/svc"
	"simple-admin-user-rpc/internal/utils/dberrorhandler"
	"simple-admin-user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.IDReq) (*user.UserInfo, error) {
	result, err := l.svcCtx.DB.User.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &user.UserInfo{
		Id:        result.ID,
		CreatedAt: result.CreatedAt.UnixMilli(),
		UpdatedAt: result.UpdatedAt.UnixMilli(),
		Name:      result.Name,
		Gender:    string(result.Gender),
		Avatar:    result.Avatar,
		Username:  result.Username,
		Mobile:    result.Mobile,
		IdCard:    result.IDCard,
		Openpid:   result.Openpid,
	}, nil
}
