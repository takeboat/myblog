package user

import (
	"context"
	"errors"
	"regexp"

	"blog/api/internal/model"
	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.BaseResp, err error) {
	// 参数校验
	if err = l.validateRegisterReq(req); err != nil {
		resp = utils.NewErrRespWithMessage(utils.InvalidParameter, err.Error())
		return
	}
	// 检查用户名是否存在
	user, err := l.svcCtx.UserModel.FindByUsername(req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	if user != nil {
		resp = utils.NewErrRespWithCode(utils.UserAlreadyExists)
		return
	}
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		resp = utils.NewErrRespWithMessage(utils.UnknownError, "密码加密失败")
		return
	}
	user = &model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		Nickname: req.Nickname,
	}
	err = l.svcCtx.UserModel.Insert(user)
	if err != nil {
		resp = utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	resp = utils.NewSuccessResp()
	return
}

func (l *RegisterLogic) validateRegisterReq(req *types.RegisterReq) (err error) {
	if req.Username == "" || req.Password == "" || req.Email == "" || req.Nickname == "" {
		return errors.New("参数不可以为空")
	}
	if len(req.Nickname) < 1 || len(req.Nickname) > 20 {
		return errors.New("昵称长度必须在1-20个字符之间")
	}
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, req.Email)
	if !matched {
		return errors.New("邮箱格式不正确")
	}
	if len(req.Email) > 100 {
		return errors.New("邮箱长度不能超过100个字符")
	}
	if len(req.Username) < 1 || len(req.Username) > 20 {
		return errors.New("用户名长度必须在1-20个字符之间")
	}
	return nil
}
