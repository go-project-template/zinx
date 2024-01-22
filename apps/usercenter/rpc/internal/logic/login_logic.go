package logic

import (
	"context"
	"zinx-zero/apps/model"
	"zinx-zero/apps/usercenter/rpc/usercenter"

	"zinx-zero/apps/acommon/aerr"
	"zinx-zero/apps/acommon/autils"
	"zinx-zero/apps/acommon/globalkey"
	"zinx-zero/apps/usercenter/rpc/internal/svc"
	"zinx-zero/apps/usercenter/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = aerr.NewErrMsg("生成token失败")
var ErrUsernamePwdError = aerr.NewErrMsg("账号或密码不正确")
var ErrUserNoExistsError = aerr.NewErrMsg("用户不存在")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {

	var accountId int64
	var err error
	switch in.AuthType {
	case globalkey.Model_UserAuthTypeSystem:
		accountId, err = l.loginByMobile(in.AuthKey, in.Password)
	default:
		return nil, aerr.NewErrCode(aerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, err
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		AccountId: accountId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken accountId : %d", accountId)
	}

	return &usercenter.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {

	user, err := l.svcCtx.UserAccountModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(aerr.NewErrCode(aerr.DB_ERROR), "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "mobile:%s", mobile)
	}

	if !(autils.Md5HexByString(password) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配出错")
	}

	return user.AccountId, nil
}

func (l *LoginLogic) loginBySmallWx() error {
	return nil
}
