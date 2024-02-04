package botMsg

import (
	"SamgeWxApi/cmd/wxBot/botUtil"
	"github.com/eatmoreapple/openwechat"
)

// 处理公众号业务

// OnUserMp 注册根据消息发送者的行为是否匹配【公众号】的消息处理函数
func OnUserMp(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnUser(checkUserMp, func(ctx *openwechat.MessageContext) {
		debugPrintMsg("OnUserMp 注册根据消息发送者的行为是否匹配【公众号】的消息处理函数", getSenderNameAndRawContent(ctx))

		msg := ctx.Message
		sender := botUtil.GetMsgSenderWithoutGroup(msg, "获取[OnUserMp]消息发送者")
		if sender == nil {
			return
		}

		botUtil.SaveMpChatLog(msg, msg.Content, "公众号消息不进行回复（后续可以用来监控公众号推文抓取~）", "公众号")
		if msg.IsArticle() {
			articles := botUtil.GetArticles(msg)
			botUtil.SaveMpArticleLog(msg, articles.String(), "公众号文章推送")
		}

	})
}

// checkUserMp 检查用户是否符合要求: 公众号
func checkUserMp(user *openwechat.User) bool {
	return user.IsMP()
}
