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
	_, err := l.svcCtx.DB.User.Delete().Where(user2.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &user.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
