package tag

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

type CreateTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTagLogic {
	return &CreateTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTagLogic) CreateTag(req *types.CreateTagReq) (resp *types.BaseResp, err error) {
	// 检查标签名称是否已存在
	_, err = l.svcCtx.TagModel.FindByName(req.Name)
	if err == nil {
		// 标签已存在，返回相应错误码
		resp = utils.NewErrRespWithCode(utils.TagAlreadyExists)
		return
	}
	// 如果查询出错且不是"记录未找到"错误，则返回数据库错误
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}

	// 创建新标签
	err = l.svcCtx.TagModel.Insert(&model.Tag{Name: req.Name})
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}

	resp = utils.NewSuccessResp()
	return
}
