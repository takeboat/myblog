package post

import (
	"context"
	"errors"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostDetailLogic {
	return &PostDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostDetailLogic) PostDetail(req *types.IdReq) (resp *types.Post, err error) {
	post, err := l.svcCtx.PostModel.GetPostWithTags(req.Id)
	if err != nil {
		err = errors.New(utils.ErrorCodeMessages[utils.DatabaseError])
		return
	}

	go l.svcCtx.PostModel.IncreaseViewCount(req.Id)
	resp = &types.Post{
		Id:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		UserId:    post.UserID,
		ViewCount: post.ViewCount + 1,
		Status:    post.Status,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	if post.Category != nil {
		resp.CategoryId = &post.Category.ID
	}
	if post.Tags != nil {
		for _, tag := range post.Tags {
			resp.Tags = append(resp.Tags, types.Tag{
				Id:        tag.ID,
				Name:      tag.Name,
				CreatedAt: tag.CreatedAt.Format("2006-01-02 15:04:05"),
			})
		}
	}
	return

}
