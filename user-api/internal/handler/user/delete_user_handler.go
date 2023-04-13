package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"simple-admin-user-api/internal/logic/user"
	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"
)

// swagger:route post /user/delete user DeleteUser
//
// Delete user information | 删除User信息
//
// Delete user information | 删除User信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDeleteUserLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUser(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
