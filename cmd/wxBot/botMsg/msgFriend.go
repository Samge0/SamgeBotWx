package botMsg

import (
	"SamgeWxApi/cmd/utils/u_str"
	"SamgeWxApi/cmd/wxBot/botConfig"
	"SamgeWxApi/cmd/wxBot/botUtil"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"os"
	"strings"
)

// 处理好友消息业务

// OnFriend 注册发送者为好友的处理函数
func OnFriend(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnFriend(func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		sender := botUtil.GetMsgSenderWithoutGroup(msg, "获取[OnFriend]消息发送者")
		if sender == nil {
			return
		}

		// 过滤公众号
		if sender.IsMP() {
			return
		}

		name := sender.NickName
		friendIds := u_str.FirstStr(botConfig.FriendIds, os.Getenv(botConfig.EnvKeyFriendIds))
		needParseMsg := friendIds == "" || strings.Contains(friendIds, name)
		if needParseMsg {
			qType := fmt.Sprintf("[好友]%s|%s|%s|%s", name, sender.ID(), sender.DisplayName, sender.RemarkName)
			botUtil.CheckStartTagAndReply(msg, qType, sender)
		}
	})
}
