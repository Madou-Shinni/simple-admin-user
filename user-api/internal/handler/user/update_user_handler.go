package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"simple-admin-user-api/internal/logic/user"
	"simple-admin-user-api/internal/svc"
	"simple-admin-user-api/internal/types"
)

// swagger:route put /user/update user UpdateUser
//
// Update user information | 更新User
//
// Update user information | 更新User
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UserInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUser(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}