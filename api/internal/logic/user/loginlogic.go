package user

import (
	"context"
	"time"

	"blog/api/internal/svc"
	"blog/api/internal/types"
	"blog/api/internal/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = &types.LoginResp{}
	user, err := l.svcCtx.UserModel.FindByEmail(req.Email)
	if err != nil {
		resp.BaseResp = *utils.NewErrRespWithCode(utils.DatabaseError)
		return
	}
	if user == nil {
		resp.BaseResp = *utils.NewErrRespWithCode(utils.UserNotFound)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		resp.BaseResp = *utils.NewErrRespWithCode(utils.InvalidCredentials)
		return
	}
	token, err := l.generateToken(user.ID, user.Username)
	if err != nil {
		resp.BaseResp = *utils.NewErrRespWithCode(utils.GenerateTokenFailed)
		return
	}
	resp.BaseResp = *utils.NewSuccessResp()

	resp.AccessToken = token
	resp.ExpiresIn = l.svcCtx.Config.Auth.AccessExpire
	return
}

func (l *LoginLogic) generateToken(userID int64, username string) (string, error) {
	claims := &jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
}
