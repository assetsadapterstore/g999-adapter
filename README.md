# g999-adapter

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建G999.ini文件，编辑如下内容：

```ini

# RPC Server Type，0: CoreWallet RPC; 1: Explorer API
rpcServerType = 0
# node api url, if RPC Server Type = 0, use bitcoin core full node
;serverAPI = "http://127.0.0.1:8333/"
# RPC Authentication Username
rpcUser = "user"
# RPC Authentication Password
rpcPassword = "password"
# Is network test?
isTestNet = false
# minimum transaction fees
minFees = "3.2"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```

