package post

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPostsLogic {
	return &ListPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPostsLogic) ListPosts(req *types.PostListReq) (resp *types.PostListResp, err error) {
	resp = &types.PostListResp{}
	page := req.Page
	pageSize := req.PageSize

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	posts, total, err := l.svcCtx.PostModel.List(page, pageSize, req.CategoryId)
	if err != nil {
		resp.BaseResp = *utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	postList := make([]types.Post, 0, len(posts))
	for _, post := range posts {
		postItem := types.Post{
			Id:        post.ID,
			Title:     post.Title,
			Content:   post.Content,
			UserId:    post.UserID,
			Status:    post.Status,
			ViewCount: post.ViewCount,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		if post.CategoryID != nil {
			postItem.CategoryId = post.CategoryID
		}
		if len(post.Tags) > 0 {
			for _, tag := range post.Tags {
				postItem.Tags = append(postItem.Tags, types.Tag{
					Id:        tag.ID,
					Name:      tag.Name,
					CreatedAt: tag.CreatedAt.Format("2006-01-02 15:04:05"),
				})
			}
		}
		postList = append(postList, postItem)
	}
	resp = &types.PostListResp{
		BaseResp: *utils.NewSuccessResp(),
		List:     postList,
		Total:    total,
	}
	return
}
