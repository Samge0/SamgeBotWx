package botConfig

import "SamgeWxApi/cmd/utils/u_file"

const (
	BotCacheDir      = "tmp/wxBotCache/botCacheFile"
	BotLogDir        = "tmp/wxBotCache/botLog"
	LoginStoragePath = "tmp/wxBotCache/storage.json"

	EnvKeyGroupIds     = "GROUP_IDS"
	EnvKeyFriendIds    = "FRIEND_IDS"
	EnvKeyOpenAiToken  = "OPEN_AI_TOKEN"
	EnvKeyMineNickname = "MINE_NICKNAME"

	GroupIds     = ""
	FriendIds    = ""
	OpenAiToken  = ""
	MineNickname = ""

	// 下面是测试数据
	//BotCacheDir      = "tmp/wxBotCache/botCacheFile"
	//BotLogDir        = "tmp/wxBotCache/botLog"
	//LoginStoragePath = "tmp/wxBotCache/storage.json"
	//GroupIds         = "Samge测试群,其他逗比群"
	//FriendIds        = "Samge,SamgeBot"
	//OpenAiToken      = "sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	//MineNickname     = "STE"
)

// InitCacheDir 初始化缓存目录
func InitCacheDir() error {
	if err := CheckAndCreateCacheDir(BotCacheDir); err != nil {
		return err
	}
	if err := CheckAndCreateCacheDir(BotLogDir); err != nil {
		return err
	}
	return nil
}

// CheckAndCreateCacheDir 检查并创建缓存目录
func CheckAndCreateCacheDir(dirPath string) error {
	if err := u_file.CreateMultiDir(dirPath); err != nil {
		return err
	}
	return nil
}
