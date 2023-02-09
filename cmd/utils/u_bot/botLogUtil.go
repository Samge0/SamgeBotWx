package u_bot

import (
	"SamgeWxApi/cmd/utils/u_date"
	"SamgeWxApi/cmd/utils/u_file"
	"fmt"
)

// SaveChatLog 保存聊天日志
func SaveChatLog(dir, sengInfo string, question string, answer string, qType string) {
	chatLogPath := fmt.Sprintf("%s/chatRecord.log", dir)
	spiltLine := "-----------------------------------------------------------------------------"
	currDate := u_date.GetCurrentDateStr()
	_format := "%s【%s】\n%s：%s\nBot：%s\n%s\n\n"
	logTxt := fmt.Sprintf(_format, currDate, qType, sengInfo, question, answer, spiltLine)
	_ = u_file.SaveAppend(chatLogPath, logTxt)
}
