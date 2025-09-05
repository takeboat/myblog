package category

import (
	"context"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCategoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCategoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCategoriesLogic {
	return &ListCategoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCategoriesLogic) ListCategories() (resp *types.CategoryListResp, err error) {
	categories, err := l.svcCtx.CategoryModel.List()
	if err != nil {
		resp = &types.CategoryListResp{}
		base := utils.NewErrRespWithCode(utils.DatabaseError)
		resp.BaseResp = *base
		return
	}
	resp = &types.CategoryListResp{
		List: make([]types.Category, len(categories)),
	}
	for i, category := range categories {
		resp.List[i] = types.Category{
			Id:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	resp.BaseResp = *utils.NewSuccessResp()
	return
}
