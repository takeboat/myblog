package tag

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTagLogic) CreateTag(req *types.CreateTagReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
