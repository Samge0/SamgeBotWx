package botUtil

import (
	"SamgeWxApi/cmd/utils/u_date"
	"SamgeWxApi/cmd/utils/u_file"
	"SamgeWxApi/cmd/utils/u_str"
	"SamgeWxApi/cmd/wxBot/botConfig"
	"SamgeWxApi/cmd/wxBot/botModel"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"strconv"
)

// SaveMpArticleLog 保存公众号推送文章信息
func SaveMpArticleLog(msg *openwechat.Message, articleLog string, qType string) {
	mpArticleLogPath := fmt.Sprintf("%s/wxMpArticle2.log", botConfig.BotLogDir)
	spiltLine := "-----------------------------------------------------------------------------"
	currDate := u_date.GetCurrentDateStr()
	_format := "%s【%s】\n%s：%s\n%s\n\n"
	sender, _ := msg.Sender()
	logTxt := fmt.Sprintf(_format, currDate, qType, sender, articleLog, spiltLine)
	_ = u_file.SaveAppend(mpArticleLogPath, logTxt)
}

// SaveMpChatLog 保存公众号聊天日志
func SaveMpChatLog(msg *openwechat.Message, question string, answer string, qType string) {
	mpChatLogPath := fmt.Sprintf("%s/wxMpChat.log", botConfig.BotLogDir)
	spiltLine := "-----------------------------------------------------------------------------"
	currDate := u_date.GetCurrentDateStr()
	_format := "%s【%s】\n%s：%s\nBot：%s\n%s\n\n"
	sender, _ := msg.Sender()
	logTxt := fmt.Sprintf(_format, currDate, qType, sender, question, answer, spiltLine)
	_ = u_file.SaveAppend(mpChatLogPath, logTxt)
}

// GetArticles 获取公众号的推文列表
func GetArticles(msg *openwechat.Message) botModel.Articles {
	if msg.IsArticle() {
		return GetArticlesByStr(msg.Content)
	}
	return make(botModel.Articles, 0)
}

// GetArticlesByStr 获取公众号的推文列表
func GetArticlesByStr(xmlStr string) botModel.Articles {
	articles := make(botModel.Articles, 0)
	var article botModel.WxMpArticleMsg
	if err := u_file.ReadXmlByStr(xmlStr, &article); err != nil {
		return articles
	}
	// 添加item推文（可能有多条）
	for _, item := range article.Appmsg.Mmreader.Category.Item {
		itemArticle := GenerateWxMpArticleItem(
			u_str.TrimNewlineSpace(item.Title),
			u_str.TrimNewlineSpace(item.URL),
			GetMpAppName(article, item),
			GetMpCoverItem(item),
			GetMpPubTimeByItem(item),
		)
		articles = append(articles, &itemArticle)
	}
	return articles
}

// GetMpPubTimeByItem 获取公众号推文时间
func GetMpPubTimeByItem(item *botModel.WxAppMsgItem) string {
	var pubTime string
	pubTime64, err := strconv.ParseInt(item.PubTime, 10, 64)
	if err != nil {
		pubTime = u_date.GetCurrentDateStr()
	} else {
		pubTime = u_date.GetTimeStampStr(pubTime64)
	}
	return pubTime
}

// GetMpPubTime 获取公众号推文时间，不推荐，因为每个item中有单独的时间，这里的时间实际也是取首个item的推送时间
func GetMpPubTime(article botModel.WxMpArticleMsg) string {
	itemLst := article.Appmsg.Mmreader.Category.Item
	if len(itemLst) == 0 {
		return u_date.GetCurrentDateStr()
	}
	return GetMpPubTimeByItem(itemLst[0])
}

// GetMpAppName 获取公众号名称
func GetMpAppName(article botModel.WxMpArticleMsg, item *botModel.WxAppMsgItem) string {
	return u_str.FirstStr(
		u_str.TrimNewlineSpace(item.Sources.Source.Name),
		u_str.TrimNewlineSpace(article.Appinfo.Appname),
		u_str.TrimNewlineSpace(article.Appmsg.Mmreader.Publisher.Nickname),
		u_str.TrimNewlineSpace(article.Appmsg.Mmreader.Category.Name),
	)
}

// GetMpCoverMain 获取公众号推文主图
func GetMpCoverMain(article botModel.WxMpArticleMsg) string {
	return u_str.FirstStr(
		u_str.TrimNewlineSpace(article.Appmsg.Thumburl),
		u_str.TrimNewlineSpace(article.Appmsg.Mmreader.Category.Topnew.Cover),
	)
}

// GetMpCoverItem 获取公众号推文Item图
func GetMpCoverItem(item *botModel.WxAppMsgItem) string {
	return u_str.FirstStr(
		u_str.TrimNewlineSpace(item.Cover),
		u_str.TrimNewlineSpace(item.Cover169),
		u_str.TrimNewlineSpace(item.Cover11),
	)
}

// GenerateWxMpArticleItem 生成一个微信公众号item
func GenerateWxMpArticleItem(
	title string,
	url string,
	source string,
	cover string,
	pubTime string,
) botModel.WxMpArticleItem {
	return botModel.WxMpArticleItem{
		Title:   u_str.TrimNewlineSpace(title),
		URL:     u_str.TrimNewlineSpace(url),
		Source:  source,
		Cover:   cover,
		PubTime: pubTime,
	}
}
