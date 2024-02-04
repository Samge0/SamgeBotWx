package botHandler

import (
	config "SamgeWxApi/cmd/utils/u_config"
	"SamgeWxApi/cmd/wxBot/botUtil"
	"github.com/eatmoreapple/openwechat"
	"io"
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
	reloadStorage := openwechat.NewFileHotReloadStorage(config.LoginStoragePath)
	defer func(reloadStorage io.ReadWriteCloser) {
		err := reloadStorage.Close()
		if err != nil {
			botUtil.SaveErrorLog(err, "reloadStorage.Close")
		}
	}(reloadStorage)
	if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		botUtil.SaveErrorLog(err, "推送登录")
		return false
	}
	return true
}
