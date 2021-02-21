package openwtester

import (
	"github.com/assetsadapterstore/g999-adapter/g999"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(g999.Symbol, g999.NewWalletManager())
}
