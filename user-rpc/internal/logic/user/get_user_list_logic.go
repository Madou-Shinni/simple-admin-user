package user

import (
	"context"
	"entgo.io/ent/dialect/sql"
	user2 "simple-admin-user-rpc/ent/user"

	"simple-admin-user-rpc/ent/predicate"
	"simple-admin-user-rpc/internal/svc"
	"simple-admin-user-rpc/internal/utils/dberrorhandler"
	"simple-admin-user-rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *user.UserListReq) (*user.UserListResp, error) {
	var predicates []predicate.User
	if in.Name != "" {
		predicates = append(predicates, user2.NameContains(in.Name))
	}
	if in.Gender != "" {
		predicates = append(predicates, sql.FieldContains(user2.FieldGender, in.Gender))
	}
	if in.Avatar != "" {
		predicates = append(predicates, user2.AvatarContains(in.Avatar))
	}
	result, err := l.svcCtx.DB.User.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &user.UserListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &user.UserInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			UpdatedAt: v.UpdatedAt.UnixMilli(),
			Name:      v.Name,
			Gender:    string(v.Gender),
			Avatar:    v.Avatar,
			Username:  v.Username,
			Mobile:    v.Mobile,
			IdCard:    v.IDCard,
			Openpid:   v.Openpid,
		})
	}

	return resp, nil
}
