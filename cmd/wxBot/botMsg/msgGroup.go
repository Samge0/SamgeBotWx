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

// 处理群组消息业务

// OnGroup 注册发送者为群组的处理函数
func OnGroup(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnGroup(func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		sender := botUtil.GetMsgSender(msg, "获取[OnGroup]消息发送者sender")
		senderInGroup := botUtil.GetMsgSenderInGroup(msg, "获取[OnGroup]消息发送者senderInGroup")

		mineNicknameInfo := fmt.Sprintf("@%s", u_str.FirstStr(botConfig.MineNickname, os.Getenv(botConfig.EnvKeyMineNickname)))
		if sender == nil || senderInGroup == nil || !strings.Contains(msg.Content, mineNicknameInfo) {
			return
		}

		name := sender.NickName
		groupIds := u_str.FirstStr(botConfig.GroupIds, os.Getenv(botConfig.EnvKeyGroupIds))
		needParseMsg := groupIds == "" || (len(name) > 0 && strings.Contains(groupIds, name))
		if needParseMsg {
			fmt.Println(msg.Content)
			qType := fmt.Sprintf("[群组]%s", name)
			botUtil.CheckStartTagAndReply(msg, qType, senderInGroup)
		}
	})
}
