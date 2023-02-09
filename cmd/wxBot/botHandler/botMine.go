package botHandler

import (
	"SamgeWxApi/cmd/wxBot/botUtil"
	"github.com/eatmoreapple/openwechat"
)

// ParseMine 处理当前登录用户相关事务
func ParseMine(bot *openwechat.Bot) (*openwechat.Self, error) {
	self, err := bot.GetCurrentUser()
	if err != nil {
		botUtil.SaveErrorLog(err, "获取登陆的用户")
		return nil, nil
	}
	return self, err
}
