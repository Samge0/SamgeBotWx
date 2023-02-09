package botHandler

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

// ParseGroups 处理群组相关事务
func ParseGroups(self *openwechat.Self) {
	groups, err := self.Groups()
	fmt.Println(groups, err)
}
