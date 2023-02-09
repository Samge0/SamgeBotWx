package botHandler

import "github.com/eatmoreapple/openwechat"

// CreatBot 创建wx机器人
func CreatBot() *openwechat.Bot {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式
	return bot
}
