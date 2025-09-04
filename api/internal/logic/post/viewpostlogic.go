package post

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewPostLogic {
	return &ViewPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewPostLogic) ViewPost(req *types.IdReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
