package botHandler

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}
