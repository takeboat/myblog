package tag

import (
	"context"
	"errors"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DeleteTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTagLogic) DeleteTag(req *types.IdReq) (resp *types.BaseResp, err error) {
	// 首先检查标签是否存在
	_, err = l.svcCtx.TagModel.FindByID(req.Id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// 标签不存在，返回相应错误码
		resp = utils.NewErrRespWithCode(utils.TagNotFound)
		return
	}

	// 执行删除操作
	err = l.svcCtx.TagModel.Delete(req.Id)
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	resp = utils.NewSuccessResp()
	return
}
