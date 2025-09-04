package user

import (
	"context"
	"fmt"
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
	user, err := l.svcCtx.UserModel.FindByUsername(req.Username)
	if err != nil {
		err = fmt.Errorf("database error: %v", err)
		return
	}
	if user == nil {
		err = fmt.Errorf("%s:%s", utils.ErrorCodeMessages[utils.UserNotFound], req.Username)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		err = fmt.Errorf("%s:%s", utils.ErrorCodeMessages[utils.InvalidCredentials], req.Username)
		return
	}
	token, err := l.generateToken(user.ID, user.Username)
	if err != nil {
		err = fmt.Errorf("%s:%s", utils.ErrorCodeMessages[utils.GenerateTokenFailed], req.Username)
		return
	}
	resp = &types.LoginResp{
		AccessToken: token,
		ExpiresIn:   time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).Unix(),
	}
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
