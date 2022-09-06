package configs

import "os"

const (
	GoodCode  = 200
	GoodMsg   = "Good! ฅ^•ﻌ•^ฅ"
	LongCode  = 233
	LongMsg   = "Long time error... (´･д･｀)"
	NopCode   = 244
	NopMsg    = "Nop （´(ｪ)｀）"
	WrongCode = 555
	WrongMsg  = "出错啦，快联系管理员！ (´･д･｀)"

	LongInfo       = "killWrong"
	Domain         = "www.codesheep.xyz"
	ContentMaxSize = 1 << 20
)

var (
	WorkPath string
)

func init() {
	here, _ := os.Getwd()
	WorkPath = here + "/work/"
}
