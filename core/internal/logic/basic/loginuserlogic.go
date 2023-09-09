package basic

import (
	"context"

	"github.com/pjimming/zeus/core/internal/svc"
	"github.com/pjimming/zeus/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(req *types.LoginUserReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
