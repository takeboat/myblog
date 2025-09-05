package post

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.IdReq) (resp *types.BaseResp, err error) {
	post, err := l.svcCtx.PostModel.FindByID(req.Id)
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	if post == nil {
		resp = utils.NewErrRespWithCode(utils.PostNotFound)
		return
	}
	err = l.svcCtx.PostModel.Delete(req.Id)
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	resp = utils.NewSuccessResp()
	return
}
