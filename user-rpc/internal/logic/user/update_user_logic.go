package user

import (
	"context"

	"simple-admin-user-rpc/internal/svc"
	"simple-admin-user-rpc/internal/utils/dberrorhandler"
	"simple-admin-user-rpc/types/user"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user.UserInfo) (*user.BaseResp, error) {
	err := l.svcCtx.DB.User.UpdateOneID(in.Id).
		SetNotEmptyName(in.Name).
		SetNotEmptyAvatar(in.Avatar).
		SetNotEmptyUsername(in.Username).
		SetNotEmptyMobile(in.Mobile).
		SetNotEmptyIDCard(in.IdCard).
		SetNotEmptyOpenpid(in.Openpid).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &user.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
