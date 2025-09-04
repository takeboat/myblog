package tag

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagsLogic {
	return &ListTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTagsLogic) ListTags() (resp *types.TagListResp, err error) {
	tags, err := l.svcCtx.TagModel.List()
	if err != nil {
		return
	}
	resp = &types.TagListResp{
		List: make([]types.Tag, len(tags)),
	}
	for i, tag := range tags {
		resp.List[i] = types.Tag{
			Id:        tag.ID,
			Name:      tag.Name,
			CreatedAt: tag.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return
}
