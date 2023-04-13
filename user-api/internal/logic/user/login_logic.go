package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"simple-admin-user-rpc/types/user"
	"time"

	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserInfo) (resp *types.BaseDataInfo, err error) {
	resp = &types.BaseDataInfo{}
	// 根据openpid查询用户
	result, err := l.svcCtx.UserRpc.GetUserByOpenpid(l.ctx, &user.OpenpidReq{Openpid: req.Openpid})
	if err != nil {
		// 用户不存在，或发生错误
		resp.Code = 100
		resp.Msg = "登陆失败：" + err.Error()
		return resp, err
	}

	// 用户存在生成jwt
	user := types.UserInfo{
		Name:     result.Name,
		Gender:   result.Gender,
		Avatar:   result.Avatar,
		Username: result.Username,
		Mobile:   result.Mobile,
		IdCard:   result.IdCard,
		Openpid:  result.Openpid,
	}

	user.Id = result.Id
	user.CreatedAt = result.CreatedAt
	user.UpdatedAt = result.UpdatedAt

	// jwt密钥
	secretKey := l.svcCtx.Config.Auth.AccessSecret
	// 过期时间 秒
	expreat := l.svcCtx.Config.Auth.AccessExpire

	token, err := NewJwtToken(secretKey, expreat, user)
	if err != nil {
		// jwt生成失败
		resp.Code = 101
		resp.Msg = "登陆失败：" + err.Error()
		return resp, err
	}

	// TODO: 生成jwt测试
	marshal, err := json.Marshal(user)

	fmt.Printf("token: %s\n,userinfo: %s", token, marshal)

	// success
	resp.Msg = "success"
	resp.Code = 0
	resp.Data = token

	return resp, nil
}

// 参考simple-admin，生成jwt
func NewJwtToken(secretKey string, seconds int64, userInfo types.UserInfo) (string, error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = now + seconds
	claims["userInfo"] = userInfo
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
