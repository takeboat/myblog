package post

import (
	"context"
	"encoding/json"

	"blog/api/internal/model"
	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostReq) (resp *types.BaseResp, err error) {
	if req.Title == "" || req.Content == "" {
		resp = utils.NewErrRespWithMessage(utils.InvalidParameter, "标题或者内容不能为空")
		return
	}
	userID := l.getUserId()
	if userID == 0 {
		resp = utils.NewErrRespWithMessage(utils.InvalidCredentials, "用户不存在")
		return
	}
	tx := l.svcCtx.PostModel.BeginTx()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	post := &model.Post{
		Title:     req.Title,
		Content:   req.Content,
		UserID:    userID,
		Status:    1,
		ViewCount: 0,
	}
	if req.CategoryId != nil {
		post.CategoryID = req.CategoryId
	}
	err = l.svcCtx.PostModel.Insert(post)
	if err != nil {
		resp = utils.NewErrRespWithMessage(utils.DatabaseError, "创建文章失败")
		return
	}
	if len(req.TagIds) > 0 {
		err = l.svcCtx.PostModel.AddTagsWithTx(tx, post.ID, req.TagIds)
		if err != nil {
			resp = utils.NewErrRespWithMessage(utils.DatabaseError, "添加标签失败")
			return
		}
	}
	resp = utils.NewSuccessResp()
	return
}

func (l *CreatePostLogic) getUserId() int64 {
	if uid, ok := l.ctx.Value("user_id").(json.Number); ok {
		if int64Uid, err := uid.Int64(); err == nil {
			return int64Uid
		}
	}
	return 0
}
