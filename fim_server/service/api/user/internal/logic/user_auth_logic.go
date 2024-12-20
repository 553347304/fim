package logic

import (
	"context"
	"fim_server/models/user_models"
	"fim_server/utils/stores/logs"

	"fim_server/service/api/user/internal/svc"
	"fim_server/service/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAuthLogic {
	return &UserAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAuthLogic) UserAuth(req *types.UserAuthRequest) (resp *types.UserAuthResponse, err error) {
	// todo: add your logic here and delete this line

	var friend user_models.FriendModel
	if friend.IsFriend(l.svcCtx.DB, req.UserId, req.FriendId) {
		return nil, logs.Error("已经是好友了")
	}

	var userConfig user_models.UserConfigModel
	err = l.svcCtx.DB.Take(&userConfig, "user_id = ?", req.FriendId).Error
	if err != nil {
		return nil, logs.Error("用户不存在")
	}
	resp = new(types.UserAuthResponse)
	resp.Verify = userConfig.Verify
	switch userConfig.Verify {
	case 0:
		return nil, logs.Error("不允许任何人添加")
	case 1:
	// 允许任何人添加
	case 2:
	// 需要验证
	case 3, 4:
		// 需要回答问题   需要正确回答问题
		resp.VerifyInfo = types.VerifyInfo{
			Issue:  userConfig.VerifyInfo.Issue,
			Answer: userConfig.VerifyInfo.Answer,
		}
	}

	return
}
