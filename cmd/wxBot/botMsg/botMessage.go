package botMsg

import (
	"SamgeWxApi/cmd/wxBot/botUtil"
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

// ParseMessage 注册消息处理函数
func ParseMessage(bot *openwechat.Bot) {
	dispatcher := openwechat.NewMessageMatchDispatcher()
	regDispatcher(dispatcher)
	bot.MessageHandler = dispatcher.AsMessageHandler()
}

// regDispatcher 注册消息调度
func regDispatcher(dispatcher *openwechat.MessageMatchDispatcher) {
	// 按对象类型区分处理：如 添加好友、群组、好友、自己、公众号、指定名称的群组/好友、自定义条件的用户
	OnFriendAdd(dispatcher)
	OnFriend(dispatcher)
	OnGroup(dispatcher)
	OnUser(dispatcher)
	OnFriendByNickName(dispatcher, "")
	OnFriendByRemarkName(dispatcher, "")
	OnGroupByGroupName(dispatcher, "")
	OnUserMp(dispatcher) // 自定义监听公众号类型消息

	// 按消息类型区分处理。目前不采用这种方式，因为不同类型可以用工具类对msg统一区分处理
	//OnText(dispatcher)
	//OnImage(dispatcher)
	//OnVoice(dispatcher)
	//OnCard(dispatcher)
	//OnMedia(dispatcher)
}

// OnText 注册处理消息类型为Text的处理函数
func OnText(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnText(func(ctx *openwechat.MessageContext) {
	})
}

// OnImage 注册处理消息类型为Image的处理函数
func OnImage(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnImage(func(ctx *openwechat.MessageContext) {
	})
}

// OnEmoticon 注册处理消息类型为Emoticon的处理函数(表情包)
func OnEmoticon(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnImage(func(ctx *openwechat.MessageContext) {
	})
}

// OnVoice 注册处理消息类型为Voice的处理函数
func OnVoice(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnVoice(func(ctx *openwechat.MessageContext) {
	})
}

// OnFriendAdd 注册处理消息类型为FriendAdd的处理函数
func OnFriendAdd(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnFriendAdd(func(ctx *openwechat.MessageContext) {
		msg := ctx.Message
		friend, err := msg.Agree("已同意好友")
		if err != nil {
			botUtil.SaveErrorLog(err, "同意好友请求")
		}
		fmt.Println(friend)
	})
}

// OnCard 注册处理消息类型为Card的处理函数
func OnCard(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnCard(func(ctx *openwechat.MessageContext) {
	})
}

// OnMedia 注册处理消息类型为Media(多媒体消息，包括但不限于APP分享、文件分享)的处理函数
func OnMedia(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnMedia(func(ctx *openwechat.MessageContext) {
	})
}

// OnUser 注册根据消息发送者的行为是否匹配的消息处理函数
func OnUser(dispatcher *openwechat.MessageMatchDispatcher) {
	dispatcher.OnUser(checkUser, func(ctx *openwechat.MessageContext) {
	})
}

// OnFriendByRemarkName 注册根据好友备注是否匹配的消息处理函数
func OnFriendByRemarkName(dispatcher *openwechat.MessageMatchDispatcher, remarkName string) {
	dispatcher.OnFriendByRemarkName(remarkName, func(ctx *openwechat.MessageContext) {
	})
}

// OnGroupByGroupName 注册根据群名是否匹配的消息处理函数
func OnGroupByGroupName(dispatcher *openwechat.MessageMatchDispatcher, groupName string) {
	dispatcher.OnGroupByGroupName(groupName, func(ctx *openwechat.MessageContext) {
	})
}

// OnFriendByNickName 注册根据好友昵称是否匹配的消息处理函数
func OnFriendByNickName(dispatcher *openwechat.MessageMatchDispatcher, nickName string) {
	dispatcher.OnFriendByNickName(nickName, func(ctx *openwechat.MessageContext) {
	})
}

// checkUser 检查用户是否符合要求
func checkUser(user *openwechat.User) bool {
	return true
}
