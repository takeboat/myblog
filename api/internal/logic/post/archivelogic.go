package post

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArchiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArchiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArchiveLogic {
	return &ArchiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArchiveLogic) Archive(req *types.ArchiveReq) (resp *types.ArchiveResp, err error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	resp = &types.ArchiveResp{}
	archives, total, err := l.svcCtx.PostModel.GetArchivedPosts(req.Page, req.PageSize)
	if err != nil {
		resp.BaseResp = *utils.NewErrRespWithCode(utils.DatabaseError)
		return 
	}
	resp.BaseResp = *utils.NewSuccessResp()
	resp.Total = total
	resp.Data = archives
	resp.Page = req.Page
	return
}
