package botUtil

import (
	config "SamgeWxApi/cmd/utils/u_config"
	"SamgeWxApi/cmd/utils/u_date"
	"SamgeWxApi/cmd/utils/u_openai"
	"SamgeWxApi/cmd/utils/u_str"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strings"
)

// GetMsgSenderWithoutGroup 获取消息发送者信息，如果是群组则返回nil
func GetMsgSenderWithoutGroup(msg *openwechat.Message, errorTip string) *openwechat.User {
	if msg.IsSendByGroup() {
		return nil // 群组消息已经单独处理，忽略
	}
	return GetMsgSender(msg, errorTip)
}

// GetMsgSender 获取消息发送者信息
func GetMsgSender(msg *openwechat.Message, errorTip string) *openwechat.User {
	sender, err := msg.Sender()
	if err != nil {
		SaveErrorLog(err, errorTip)
		return nil
	}
	return sender
}

// GetMsgSenderNickName 获取消息发送者名称， 获取失败则返回空字符串
func GetMsgSenderNickName(msg *openwechat.Message) string {
	sender, err := msg.Sender()
	if err != nil {
		return ""
	}
	return sender.NickName
}

// GetMsgSenderInGroup 获取消息在群里面的发送者
func GetMsgSenderInGroup(msg *openwechat.Message, errorTip string) *openwechat.User {
	sender, err := msg.SenderInGroup()
	if err != nil {
		SaveErrorLog(err, errorTip)
		return nil
	}
	return sender
}

// GetMsgReceiver 获取消息接收者信息
func GetMsgReceiver(msg *openwechat.Message, errorTip string) *openwechat.User {
	receiver, err := msg.Receiver()
	if err != nil {
		SaveErrorLog(err, errorTip)
		return nil
	}
	return receiver
}

// ReplyWithOpenAi 使用openai的api进行回复
func ReplyWithOpenAi(msg *openwechat.Message, question string, qType string, sender *openwechat.User) {
	answer, err := u_openai.GetChatResponseWithToken(question)
	if err != nil {
		SaveErrorLog(err, "ReplyWithOpenAi")
	} else {
		ReplyText(msg, answer, sender)
		SaveChatLog(msg, question, answer, qType)
	}
}

// CheckStartTagAndReply 检查内容的起始标签，如果符合则进行回复
func CheckStartTagAndReply(msg *openwechat.Message, qType string, sender *openwechat.User) {
	switch {
	case msg.IsTickledMe():
		ParseMsgOnTickled(msg, fmt.Sprintf("%s 拍一拍", qType), sender)
	case msg.IsText(): // 文本
		ParseMsgOnText(msg, qType, sender)
	case msg.IsPicture(): // 图片
		ParseMsgOnImage(msg, qType, sender)
	case msg.IsVoice(): // 语音
		ParseMsgOnVoice(msg, qType, sender)
	case msg.IsCard(): // 卡片
		ParseMsgOnCard(msg, qType, sender)
	case msg.IsVideo(): // 视频
		ParseMsgOnVideo(msg, qType, sender)
	case msg.IsEmoticon(): // 表情包
		ParseMsgOnEmoticon(msg, qType, sender)
	case msg.IsRealtimeLocation(): // 实时位置共享
		ParseMsgOnRealtimeLocation(msg, qType, sender)
	case msg.IsLocation(): // 位置
		ParseMsgOnLocation(msg, qType, sender)
	case msg.IsTransferAccounts(): // 微信转账
		ParseMsgOnTransferAccounts(msg, qType, sender)
	case msg.IsSendRedPacket(): // 微信红包-发出
		ParseMsgOnSendRedPacket(msg, qType, sender)
	case msg.IsReceiveRedPacket(): // 微信红包-收到
		ParseMsgOnReceiveRedPacket(msg, qType, sender)
	case msg.IsRenameGroup(): // 群组重命名
		ParseMsgOnRenameGroup(msg, qType, sender)
	case msg.IsArticle(): // 文章
		ParseMsgOnArticle(msg, qType, sender)
	case msg.IsVoipInvite(): // 语音/视频邀请
		ParseMsgOnVoipInvite(msg, qType, sender)
	case msg.IsMedia(): // Media(多媒体消息，包括但不限于APP分享、文件分享
		ParseMsgOnMedia(msg, qType, sender)
	default:
		fmt.Printf("CheckStartTagAndReply 没有命中的类型，跳过")
	}

}

// ParseMsgOnTickled 处理【拍一拍】类型的消息
func ParseMsgOnTickled(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "再拍我我就把你吃了"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, msg.Content, answer, qType)
}

// ParseMsgOnText 处理【OnText】类型的消息
func ParseMsgOnText(msg *openwechat.Message, qType string, sender *openwechat.User) {
	var question string
	question = msg.Content

	// 解析管理员的特殊指令
	if config.NeedParseCmd(sender.NickName, question) {
		fmt.Printf("解析管理员(%s)的特殊指令：%s\n", sender.NickName, question)
		switch question {
		case config.CmdOpenReply:
			config.LoadConfig().EnableReply = true
		case config.CmdCloseReply:
			config.LoadConfig().EnableReply = false
		default:
			fmt.Printf("解析管理员的特殊指令: 没有命中的类型，跳过")
		}
		return
	}

	// 禁止回复则跳过
	if !config.LoadConfig().EnableReply {
		fmt.Printf("当前已禁止回复，忽略该条消息 -> %s：%s\n", sender.NickName, question)
		return
	}

	// 解析其他指令
	tagHead := "生成头像 "
	tagT2I := "生成图片 "
	switch {
	case strings.HasPrefix(question, tagHead):
		// 生成头像
		question = strings.Replace(question, tagHead, "", 1)
		answer := fmt.Sprintf("%s 服务正在开发中", tagHead)
		ReplyText(msg, answer, sender)
		SaveChatLog(msg, question, answer, tagHead)
	case strings.HasPrefix(question, tagT2I):
		// 生成图片
		question = strings.Replace(question, tagT2I, "", 1)
		answer := fmt.Sprintf("%s 服务正在开发中", tagT2I)
		ReplyText(msg, answer, sender)
		SaveChatLog(msg, question, answer, tagT2I)
	default:
		ReplyWithOpenAi(msg, question, qType, sender)
	}
}

// ParseMsgOnImage 处理【OnImage】类型的消息
func ParseMsgOnImage(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "这是什么图片"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 图片", qType))

	fileName := u_date.GetCurrentDateStr(u_date.DateFormat.Flow)
	err := msg.SaveFileToLocal(fmt.Sprintf("%s/%s_%s.jpg", config.BotCacheDir, sender.NickName, u_str.FirstStr(msg.FileName, fileName)))
	if err != nil {
		SaveErrorLog(err, "SaveFileToLocal-保存图片")
		return
	}
}

// ParseMsgOnVoice 处理【OnVoice】类型的消息
func ParseMsgOnVoice(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "不方便接听语音信息，还是发文字吧"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 语音", qType))

	fileName := u_date.GetCurrentDateStr(u_date.DateFormat.Flow)
	err := msg.SaveFileToLocal(fmt.Sprintf("%s/%s_%s.amr", config.BotCacheDir, sender.NickName, u_str.FirstStr(msg.FileName, fileName)))
	if err != nil {
		SaveErrorLog(err, "SaveFileToLocal-保存语音")
		return
	}
}

// ParseMsgOnCard 处理【OnCard】类型的消息
func ParseMsgOnCard(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "这是什么好玩的小卡片"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 卡片", qType))
}

// ParseMsgOnMedia 处理【Media(多媒体消息，包括但不限于APP分享、文件分享)的处理函数】类型的消息
func ParseMsgOnMedia(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "你这发的是个啥子哦"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 多媒体消息等", qType))

	fileName := u_date.GetCurrentDateStr(u_date.DateFormat.Flow)
	err := msg.SaveFileToLocal(fmt.Sprintf("%s/%s_%s.file", config.BotCacheDir, sender.NickName, u_str.FirstStr(msg.FileName, fileName)))
	if err != nil {
		SaveErrorLog(err, "SaveFileToLocal-保存多媒体")
		return
	}
}

// ParseMsgOnVideo 处理【视频】类型的消息
func ParseMsgOnVideo(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "这是什么视频，好看的话多发点来look look"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 视频", qType))

	fileName := u_date.GetCurrentDateStr(u_date.DateFormat.Flow)
	err := msg.SaveFileToLocal(fmt.Sprintf("%s/%s_%s.mp4", config.BotCacheDir, sender.NickName, u_str.FirstStr(msg.FileName, fileName)))
	if err != nil {
		SaveErrorLog(err, "SaveFileToLocal-保存视频")
		return
	}
}

// ParseMsgOnEmoticon 处理【表情】类型的消息
func ParseMsgOnEmoticon(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := fmt.Sprintf("%s%s", openwechat.Emoji.Awesome, openwechat.Emoji.Doge)
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 表情", qType))

	fileName := u_date.GetCurrentDateStr(u_date.DateFormat.Flow)
	err := msg.SaveFileToLocal(fmt.Sprintf("%s/%s_%s.gif", config.BotCacheDir, sender.NickName, u_str.FirstStr(msg.FileName, fileName)))
	if err != nil {
		SaveErrorLog(err, "SaveFileToLocal-保存表情")
		return
	}
}

// ParseMsgOnRealtimeLocation 处理【实时位置】类型的消息
func ParseMsgOnRealtimeLocation(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "你现在在哪个位置？我点不开看不到"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 实时位置", qType))
}

// ParseMsgOnLocation 处理【位置】类型的消息
func ParseMsgOnLocation(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "这是哪里？你又到哪里鬼混去了"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 位置", qType))
}

// ParseMsgOnTransferAccounts 处理【微信转账】类型的消息
func ParseMsgOnTransferAccounts(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "多谢老板，祝老板身体健康发大财~"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 微信转账", qType))
}

// ParseMsgOnSendRedPacket 处理【微信红包-发出】类型的消息
func ParseMsgOnSendRedPacket(msg *openwechat.Message, qType string, sender *openwechat.User) {
}

// ParseMsgOnReceiveRedPacket 处理【微信红包-收到】类型的消息
func ParseMsgOnReceiveRedPacket(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "多谢老板的大红包，好事成双，再发一个吧~"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 微信红包-收到", qType))
}

// ParseMsgOnRenameGroup 处理【群组重命名】类型的消息
func ParseMsgOnRenameGroup(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "群名变来变去的累不累哦？"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 群组重命名", qType))
}

// ParseMsgOnArticle 处理【文章消息】类型的消息
func ParseMsgOnArticle(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "这是什么绝世好文"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 文章", qType))
}

// ParseMsgOnVoipInvite 处理【语音或视频通话邀请】类型的消息
func ParseMsgOnVoipInvite(msg *openwechat.Message, qType string, sender *openwechat.User) {
	answer := "我现在在忙，不方便接听"
	ReplyText(msg, answer, sender)
	SaveChatLog(msg, "", answer, fmt.Sprintf("%s 通话邀请", qType))
}

// ReplyText 回复文本，如果是群聊，则@对方
func ReplyText(msg *openwechat.Message, answer string, sender *openwechat.User) {
	var text = answer
	var replyTypeTag = "好友"

	if msg.IsSendByGroup() && sender != nil {
		s := " "
		text = fmt.Sprintf("@%s%s%s", sender.NickName, s, answer)
		replyTypeTag = "群组"
	}

	if config.LoadConfig().EnableReply {
		_, err := msg.ReplyText(text)
		if err != nil {
			fmt.Printf("【Error】ReplyText->回复%s：%s\n", replyTypeTag, text)
		}
	} else {
		fmt.Printf("ReplyText->回复%s：%s\n", replyTypeTag, text)
	}
}
