package botUtil

import (
	config "SamgeWxApi/cmd/utils/u_config"
	"SamgeWxApi/cmd/utils/u_date"
	"SamgeWxApi/cmd/utils/u_file"
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

// SaveChatLog 保存聊天日志
func SaveChatLog(msg *openwechat.Message, question string, answer string, qType string) {
	chatLogPath := fmt.Sprintf("%s/wxChatRecord.log", config.BotLogDir)
	spiltLine := "-----------------------------------------------------------------------------"
	currDate := u_date.GetCurrentDateStr()
	_format := "%s【%s】\n%s：%s\nBot：%s\n%s\n\n"
	sender, _ := msg.Sender()
	logTxt := fmt.Sprintf(_format, currDate, qType, sender, question, answer, spiltLine)
	_ = u_file.SaveAppend(chatLogPath, logTxt)
}
