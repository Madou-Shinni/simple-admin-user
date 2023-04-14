package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"simple-admin-user-api/internal/logic/user"
	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"
)

// swagger:route get /user user GetUserById
//
// Get user by ID | 通过ID获取User
//
// Get user by ID | 通过ID获取User
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UserInfoResp

func GetUserByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUserById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
