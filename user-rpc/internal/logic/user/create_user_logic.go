package user

import (
	"context"
	user2 "simple-admin-user-rpc/ent/user"

	"simple-admin-user-rpc/internal/svc"
	"simple-admin-user-rpc/internal/utils/dberrorhandler"
	"simple-admin-user-rpc/types/user"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.UserInfo) (*user.BaseIDResp, error) {
	result, err := l.svcCtx.DB.User.Create().
		SetName(in.Name).
		SetGender(user2.Gender(in.Gender)).
		SetAvatar(in.Avatar).
		SetUsername(in.Username).
		SetMobile(in.Mobile).
		SetIDCard(in.IdCard).
		SetOpenpid(in.Openpid).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &user.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
