package user

import (
	"net/http"

	"github.com/pjimming/zeus/athana/internal/logic/user"
	"github.com/pjimming/zeus/athana/internal/svc"
	"github.com/pjimming/zeus/athana/internal/types"
	"github.com/pjimming/zeus/utils/errorx"
	"github.com/pjimming/zeus/utils/httpresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := user.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
