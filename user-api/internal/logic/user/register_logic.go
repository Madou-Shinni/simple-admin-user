package user

import (
	"context"
	"errors"
	"simple-admin-user-rpc/types/user"

	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	// 查询用户
	_, err = l.svcCtx.UserRpc.GetUserByOpenpid(l.ctx,
		&user.OpenpidReq{
			Openpid: req.Openpid,
		})
	if err == nil {
		// 有用户存在
		logx.Info("user already register")
		return nil, errors.New("用户已注册！")
	}

	// 没有用户存在，创建用户
	data, err := l.svcCtx.UserRpc.CreateUser(l.ctx,
		&user.UserInfo{
			Name:     req.Name,
			Gender:   req.Gender,
			Avatar:   req.Avatar,
			Username: req.Username,
			Mobile:   req.Mobile,
			IdCard:   req.IdCard,
			Openpid:  req.Openpid,
		})
	if err != nil {
		logx.Errorf("user-rpc create user error: %s", err.Error())
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
