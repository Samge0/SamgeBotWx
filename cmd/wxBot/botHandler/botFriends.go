package botHandler

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

// ParseFriends 处理好友相关事务
func ParseFriends(self *openwechat.Self) {
	friends, err := self.Friends()
	fmt.Println(friends, err)
}
