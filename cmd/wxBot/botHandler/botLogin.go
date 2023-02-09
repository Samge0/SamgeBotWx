package botHandler

import (
	"SamgeWxApi/cmd/wxBot/botConfig"
	"SamgeWxApi/cmd/wxBot/botUtil"
	"github.com/eatmoreapple/openwechat"
)

func ParseLogin(bot *openwechat.Bot) bool {
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	//bot.UUIDCallback = ConsoleQrCode

	// 登陆
	//if err := bot.Login(); err != nil {
	//	SaveErrorLog(err.Error(), "登录")
	//	return false
	//}

	// 热登录
	//reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	//defer reloadStorage.Close()
	//if err := bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
	//	SaveErrorLog(err.Error(), "热登录")
	//	return false
	//}

	// 推送登录
	reloadStorage := openwechat.NewFileHotReloadStorage(botConfig.LoginStoragePath)
	defer reloadStorage.Close()
	if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		botUtil.SaveErrorLog(err, "热登录")
		return false
	}
	return true
}
