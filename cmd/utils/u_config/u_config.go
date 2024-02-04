package config

import (
	"SamgeWxApi/cmd/utils/u_file"
	"SamgeWxApi/cmd/utils/u_str"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

const (
	BotCacheDir      = "tmp/wxBotCache/botCacheFile"
	BotLogDir        = "tmp/wxBotCache/botLog"
	LoginStoragePath = "tmp/wxBotCache/storage.json"

	DefaultBaseUrl = "https://api.openai.com/v1"
	DefaultModel   = "gpt-3.5-turbo-1106"

	CmdOpenReply  = "开启回复"
	CmdCloseReply = "关闭回复"
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

// IsManagerUser 是否管理员微信
func IsManagerUser(name string) bool {
	return u_str.Contains(LoadConfig().ManagerIds, name)
}

// IsCmdContent 是否指令内容
func IsCmdContent(content string) bool {
	return CmdOpenReply == content || CmdCloseReply == content
}

// NeedParseCmd 是否需要解析指令
func NeedParseCmd(name string, content string) bool {
	return IsManagerUser(name) && IsCmdContent(content)
}

// Configuration 项目配置
type Configuration struct {
	BaseUrl string `json:"base_url"` // openai的请求地址，需要携带v1版本号，默认是：${DefaultBaseUrl} ，可配置转发地址/本地部署的模型api地址
	BotDesc string `json:"bot_desc"` // 机器人引导描述词

	ApiKey           string  `json:"api_key"`    // gpt apikey
	Model            string  `json:"model"`      // GPT模型
	MaxTokens        int     `json:"max_tokens"` // GPT请求最大字符数
	Temperature      float32 `json:"temperature"`
	TopP             float32 `json:"top_p"`
	PresencePenalty  float32 `json:"presence_penalty"`
	FrequencyPenalty float32 `json:"frequency_penalty"`

	GroupIds     string `json:"group_ids"`
	FriendIds    string `json:"friend_ids"`
	ManagerIds   string `json:"manager_ids"`
	MineNickname string `json:"mine_nickname"`
	EnableReply  bool   `json:"enable_reply"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 给配置赋默认值
		config = &Configuration{
			BaseUrl: DefaultBaseUrl,
			BotDesc: "",

			ApiKey:           "",
			Model:            DefaultModel,
			MaxTokens:        60,
			Temperature:      0.9,
			TopP:             1,
			FrequencyPenalty: 0.0,
			PresencePenalty:  0.6,

			GroupIds:     "",
			FriendIds:    "",
			ManagerIds:   "",
			MineNickname: "",
			EnableReply:  true,
		}

		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				log.Fatalf("open config err: %v", err)
				return
			}
			defer func(f *os.File) {
				_ = f.Close()
			}(f)
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				log.Fatalf("decode config err: %v", err)
				return
			}
		}

		// 如果存在环境变量，则优先使用使用环境变量

		BaseUrl := os.Getenv("sg.samge_wx_bot.base_url")
		if BaseUrl != "" {
			config.BaseUrl = BaseUrl
		}

		BotDesc := os.Getenv("sg.samge_wx_bot.bot_desc")
		if BotDesc != "" {
			config.BotDesc = BotDesc
		}

		ApiKey := os.Getenv("sg.samge_wx_bot.api_key")
		if ApiKey != "" {
			config.ApiKey = ApiKey
		}

		Model := os.Getenv("sg.samge_wx_bot.model")
		if Model != "" {
			config.Model = Model
		}

		MaxTokens := os.Getenv("sg.samge_wx_bot.max_tokens")
		if MaxTokens != "" {
			max, err := strconv.Atoi(MaxTokens)
			if err != nil {
				log.Fatalf(fmt.Sprintf("config MaxTokens err: %v ,get is %v", err, MaxTokens))
				return
			}
			config.MaxTokens = max
		}

		Temperature := os.Getenv("sg.samge_wx_bot.tempreature")
		if Temperature != "" {
			temp, err := strconv.ParseFloat(Temperature, 32)
			if err != nil {
				log.Fatalf(fmt.Sprintf("config Temperature err: %v ,get is %v", err, Temperature))
				return
			}
			config.Temperature = float32(temp)
		}

		TopP := os.Getenv("sg.samge_wx_bot.top_p")
		if TopP != "" {
			temp, err := strconv.ParseFloat(TopP, 32)
			if err != nil {
				log.Fatalf(fmt.Sprintf("config Temperature err: %v ,get is %v", err, TopP))
				return
			}
			config.TopP = float32(temp)
		}

		FrequencyPenalty := os.Getenv("sg.samge_wx_bot.freq")
		if FrequencyPenalty != "" {
			temp, err := strconv.ParseFloat(FrequencyPenalty, 32)
			if err != nil {
				log.Fatalf(fmt.Sprintf("config Temperature err: %v ,get is %v", err, FrequencyPenalty))
				return
			}
			config.FrequencyPenalty = float32(temp)
		}

		PresencePenalty := os.Getenv("sg.samge_wx_bot.pres")
		if PresencePenalty != "" {
			temp, err := strconv.ParseFloat(PresencePenalty, 32)
			if err != nil {
				log.Fatalf(fmt.Sprintf("config Temperature err: %v ,get is %v", err, PresencePenalty))
				return
			}
			config.PresencePenalty = float32(temp)
		}

		GroupIds := os.Getenv("sg.samge_wx_bot.group_ids")
		if GroupIds != "" {
			config.GroupIds = GroupIds
		}

		FriendIds := os.Getenv("sg.samge_wx_bot.friend_ids")
		if FriendIds != "" {
			config.FriendIds = FriendIds
		}

		ManagerIds := os.Getenv("sg.samge_wx_bot.manager_ids")
		if ManagerIds != "" {
			config.ManagerIds = ManagerIds
		}

		MineNickname := os.Getenv("sg.samge_wx_bot.mine_nickname")
		if MineNickname != "" {
			config.MineNickname = MineNickname
		}

		EnableReply := os.Getenv("sg.samge_wx_bot.enable_reply")
		if EnableReply != "" {
			config.EnableReply = EnableReply == "true"
		}

		fmt.Printf("config.EnableReply=%v\n", config.EnableReply)
	})

	return config
}
