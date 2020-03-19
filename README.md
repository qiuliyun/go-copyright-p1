#  项目指引


## 了解项目请先阅读项目说明书
## 了解项目相关技术信息请阅读INFO.md
## 以太坊智能合约请参考solidity.md
## 本软件配置运行过程太多复杂，不再讨论

### 创建 一个本地私链

1. 新建账户
`geth --datadir $HOME/deepeth account new`
2. 编写 genesis.json

3. 初始化
`geth --datadir $HOME/deepeth init $HOME/deepeth/genesis.json` 
4. 运行
`geth --datadir $HOME/deepeth --rpc --rpcaddr 0.0.0.0 --rpcapi web3,eth,personal --ws --wsaddr 0.0.0.0 --wsorigins "*"`