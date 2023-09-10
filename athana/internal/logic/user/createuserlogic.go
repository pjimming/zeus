package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pjimming/zeus/athana/internal/svc"
	"github.com/pjimming/zeus/athana/internal/types"
	"github.com/pjimming/zeus/athana/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.UserResp, err error) {

	user := &models.User{}
	_ = copier.Copy(&user, req)
	return
}
