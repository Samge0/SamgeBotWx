package botModel

import "fmt"

// Articles 公众号文章列表
type Articles []*WxMpArticleItem

// String Articles的string输出
func (articles Articles) String() string {
	var logStr string
	for _, article := range articles {
		logStr = fmt.Sprintf("%s\n%s\n",
			logStr,
			article.String(),
		)
	}
	return logStr
}
