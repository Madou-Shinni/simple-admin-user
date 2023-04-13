package user

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"simple-admin-user-rpc/ent/predicate"
	user2 "simple-admin-user-rpc/ent/user"
	"simple-admin-user-rpc/internal/svc"
	"simple-admin-user-rpc/internal/utils/dberrorhandler"
	"simple-admin-user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByOpenPidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByOpenPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetGetUserByOpenPidLogic(in *user.OpenpidReq) (*user.UserInfo, error) {
	var predicates []predicate.User
	if in.Openpid != "" {
		predicates = append(predicates, sql.FieldEQ(user2.FieldOpenpid, in.Openpid))
	}

	result, err := l.svcCtx.DB.User.Query().Where(predicates...).First(context.Background())
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
