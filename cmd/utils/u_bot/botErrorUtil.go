package u_bot

import (
	"SamgeWxApi/cmd/utils/u_date"
	"SamgeWxApi/cmd/utils/u_file"
	"fmt"
	"log"
)

// SaveErrorLog 保存错误日志
func SaveErrorLog(dir string, content any, contentType string) {
	errorLogPath := fmt.Sprintf("%s/botError.log", dir)
	log.Println(content)
	spiltLine := "-----------------------------------------------------------------------------"
	currDate := u_date.GetCurrentDateStr()
	logTxt := fmt.Sprintf("%s：【%s】%v\n%s\n\n", currDate, contentType, content, spiltLine)
	_ = u_file.SaveAppend(errorLogPath, logTxt)
}
