package botMsg

import (
	config "SamgeWxApi/cmd/utils/u_config"
	"SamgeWxApi/cmd/utils/u_str"
	"SamgeWxApi/cmd/wxBot/botUtil"
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

// 处理好友消息业务

// OnFriend 注册发送者为好友的处理函数
func OnFriend(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnFriend(func(ctx *openwechat.MessageContext) {

		msg := ctx.Message
		sender := botUtil.GetMsgSenderWithoutGroup(msg, "获取[OnFriend]消息发送者")
		if sender == nil {
			debugPrintMsg("OnFriend 注册发送者为好友的处理函数", "sender == nil")
			return
		}

		// 过滤公众号
		if sender.IsMP() {
			return
		}

		debugPrintMsg("OnFriend 注册发送者为好友的处理函数", getSenderNameAndRawContent(ctx))

		name := sender.NickName
		friendIds := config.LoadConfig().FriendIds
		needParseMsg := friendIds == "" || u_str.Contains(friendIds, name)
		if needParseMsg {
			qType := fmt.Sprintf("[好友]%s|%s|%s|%s", name, sender.Self().ID(), sender.DisplayName, sender.RemarkName)
			botUtil.CheckStartTagAndReply(msg, qType, sender)
		} else {
			debugPrintMsg("OnFriend 注册发送者为好友的处理函数", fmt.Sprintf("%s 不在 %s 名单内，跳过消息处理", name, friendIds))
		}
	})
}
