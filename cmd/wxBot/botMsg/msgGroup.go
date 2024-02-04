package botMsg

import (
	config "SamgeWxApi/cmd/utils/u_config"
	"SamgeWxApi/cmd/utils/u_str"
	"SamgeWxApi/cmd/wxBot/botUtil"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strings"
)

// 处理群组消息业务

// OnGroup 注册发送者为群组的处理函数
func OnGroup(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnGroup(func(ctx *openwechat.MessageContext) {
		debugPrintMsg("OnGroup 注册发送者为群组的处理函数", getSenderNameAndRawContent(ctx))

		msg := ctx.Message
		sender := botUtil.GetMsgSender(msg, "获取[OnGroup]消息发送者sender")
		senderInGroup := botUtil.GetMsgSenderInGroup(msg, "获取[OnGroup]消息发送者senderInGroup")

		mineNicknameInfo := fmt.Sprintf("@%s", config.LoadConfig().MineNickname)
		if sender == nil || senderInGroup == nil || !strings.Contains(msg.Content, mineNicknameInfo) {
			return
		}

		name := sender.NickName
		groupIds := config.LoadConfig().GroupIds
		needParseMsg := groupIds == "" || (len(name) > 0 && u_str.Contains(groupIds, name))
		if needParseMsg {
			fmt.Println(msg.Content)
			qType := fmt.Sprintf("[群组]%s", name)
			botUtil.CheckStartTagAndReply(msg, qType, senderInGroup)
		}
	})
}
