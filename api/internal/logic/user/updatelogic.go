package user

import (
	"context"
	"errors"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.RegisterReq) (resp *types.BaseResp, err error) {
	user, err := l.svcCtx.UserModel.FindByUsername(req.Username)
	// 数据库错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	// 用户不存在
	if user == nil {
		resp = utils.NewErrRespWithCode(utils.UserNotFound)
		return
	}
	user.Nickname = req.Nickname
	user.Email = req.Email
	err = l.svcCtx.UserModel.Update(user)
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	resp = utils.NewRespWithMessage(utils.SuccessCode, "更新成功")
	return
}
