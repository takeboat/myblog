package post

import (
	"context"
	"encoding/json"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePostLogic) UpdatePost(req *types.UpdatePostReq) (resp *types.BaseResp, err error) {
	post, err := l.svcCtx.PostModel.FindByID(req.Id)
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	if post.UserID != l.getCurrentUserID() {
		resp = utils.NewErrRespWithMessage(utils.InvalidCredentials, "用户无权限修改该文章")
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
	if tx.Error != nil {
		resp = utils.NewErrRespWithMessage(utils.DatabaseError, "开启事务失败")
		return
	}
	// 更新文章
	post.Title = req.Title
	post.Content = req.Content
	if req.CategoryId != nil {
		post.CategoryID = req.CategoryId
	}
	err = l.svcCtx.PostModel.UpdateWithTx(tx, post)
	if err != nil {
		resp = utils.NewErrRespWithMessage(utils.DatabaseError, "更新文章失败")
		return
	}
	// 删除旧标签 添加新标签
	if len(req.TagIds) > 0 {
		err = l.svcCtx.PostModel.RemoveTagsWithTx(tx, post.ID)
		if err != nil {
			resp = utils.NewErrRespWithMessage(utils.DatabaseError, "删除标签失败")
			return
		}
		err = l.svcCtx.PostModel.AddTagsWithTx(tx, post.ID, req.TagIds)
		if err != nil {
			resp = utils.NewErrRespWithMessage(utils.DatabaseError, "添加标签失败")
			return
		}
	}
	resp = utils.NewSuccessResp()
	return
}

func (l *UpdatePostLogic) getCurrentUserID() int64 {
	// 从context中获取user_id
	// 根据go-zero框架机制，JWT claims中的"user_id"会被放到context中
	if uid, ok := l.ctx.Value("user_id").(json.Number); ok {
		if int64Uid, err := uid.Int64(); err == nil {
			return int64Uid
		}
	}

	return 0
}
