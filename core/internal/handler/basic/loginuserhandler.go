package basic

import (
	"net/http"

	"github.com/pjimming/zeus/core/internal/logic/basic"
	"github.com/pjimming/zeus/core/internal/svc"
	"github.com/pjimming/zeus/core/internal/types"
	"github.com/pjimming/zeus/utils/errorx"
	"github.com/pjimming/zeus/utils/httpresp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpresp.HttpError(w, r, errorx.NewCodeError(2, http.StatusBadRequest, err.Error()))
			return
		}

		l := basic.NewLoginUserLogic(r.Context(), svcCtx)
		resp, err := l.LoginUser(&req)

		httpresp.HTTP(w, r, resp, err)

	}
}
