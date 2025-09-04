package category

import (
	"context"
	"errors"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryLogic {
	return &DeleteCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryLogic) DeleteCategory(req *types.IdReq) (resp *types.BaseResp, err error) {
	// 查询数据库中是否存在分类
	_, err = l.svcCtx.CategoryModel.FindByID(req.Id)

	// 如果找不到这个分类 返回错误
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		resp = utils.NewErrRespWithCode(utils.CategoryNotFound)
		return
	}
	// 删除
	err = l.svcCtx.CategoryModel.Delete(req.Id)
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	resp = utils.NewSuccessResp()
	return
}
