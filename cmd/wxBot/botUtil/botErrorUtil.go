package botUtil

import (
	config "SamgeWxApi/cmd/utils/u_config"
	"SamgeWxApi/cmd/utils/u_date"
	"SamgeWxApi/cmd/utils/u_file"
	"fmt"
	"log"
)

// SaveErrorLog 保存错误日志
func SaveErrorLog(content any, contentType string) {
	errorLogPath := fmt.Sprintf("%s/wxRobotError.log", config.BotLogDir)
	log.Println(content)
	spiltLine := "-----------------------------------------------------------------------------"
	currDate := u_date.GetCurrentDateStr()
	logTxt := fmt.Sprintf("%s：【%s】%v\n%s\n\n", currDate, contentType, content, spiltLine)
	_ = u_file.SaveAppend(errorLogPath, logTxt)
}
