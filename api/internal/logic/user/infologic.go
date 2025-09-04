package user

import (
	"context"
	"fmt"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *InfoLogic) Info(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindByUsername(req.Username)
	if err != nil {
		return
	}
	if user == nil {
		err = fmt.Errorf("%s: %s", utils.ErrorCodeMessages[utils.UserNotFound], req.Username)
		return
	}
	resp = &types.UserInfoResp{
		Id:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return
}
