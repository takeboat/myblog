package category

import (
	"context"
	"errors"

	"blog/api/internal/model"
	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategoryReq) (resp *types.BaseResp, err error) {
	// 查询数据库中是否存在分类
	_, err = l.svcCtx.CategoryModel.FindByName(req.Name)
	if err == nil {
		resp = utils.NewErrRespWithCode(utils.CategoryAlreadyExists)
		return
	}
	// 如果不是没有这条记录的错误那就记录为数据库错误
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	// 插入新的分类
	err = l.svcCtx.CategoryModel.Insert(&model.Category{Name: req.Name})
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	resp = utils.NewSuccessResp()
	return
}
