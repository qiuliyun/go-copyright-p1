package eths

import (
	"time"
	"strconv"
	"math/big"
	"github.com/ethereum/go-ethereum/rpc"
	"fmt"
	"os"
	"go-copyright-p1/configs"
	"go-copyright-p1/utils"
	"github.com/onrik/ethrpc"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

//创建账户
func NewAcc(pass,connstr string)(string,error){
	//连接geth
	client,err :=rpc.Dial(connstr)
	if err !=nil{
		fmt.Println("faild to con to geth",err)
		return "",err
	}
	//创建账户
	var account string
	err =client.Call(&account,"personal_newAccount",pass)
	if err !=nil{
		fmt.Println("faild to create acc",err)
		return "",err
	}	
	fmt.Println("acc=",account)
	cli := ethrpc.New("http://localhost:8545")
	_, err = cli.Web3ClientVersion()
	if err !=nil{
		fmt.Println("失败 to web3",err)
		return "",err
	}
	//转账10eth
	var i int
	for i=0;i<10;i++{
		_,err =cli.EthSendTransaction(ethrpc.T{ From: "0xd4695dcef59587c4bc70c0da32abd2bcf9ad8ede", To: account, Value: ethrpc.Eth1(), })
		if err !=nil{
		fmt.Println("失败 to tran",err)
		return "",err
		}
	}
	return account,nil
}

//上传资产（挖矿）

func UploadPic(from,pass,hash,data string) error{
		//建立连接
	cli,err :=ethclient.Dial(configs.Config.Eth.Connstr)
	if err!= nil{
		fmt.Println("faild to geth",err)
		return err
	}
	//函数入口
	instance,err :=NewEths(common.HexToAddress(configs.Config.Eth.QaAddr), cli)
	if err!= nil{
		fmt.Println("faild to NewCallabi",err)
		return err
	}
	//获取keystore文件名
	fileName,err := utils.GetFileName(string([]rune(from)[2:]),configs.Config.Eth.Keydir)
	if err !=nil{
		fmt.Println("faild to GetFileName",err)
		return err
	}
	//设置签名
	file,err :=os.Open(configs.Config.Eth.Keydir+"/"+fileName)
	if err!= nil{
		fmt.Println("faild to Open",err)
		return err
	}
	
	auth,err :=bind.NewTransactor(file,pass)
	if err!= nil{
		fmt.Println("faild to bind",err)
		return err
	}
	//
	
	_,err =instance.Mint(auth,common.HexToHash(hash),data)
		if err!= nil{
		fmt.Println("faild to .Mint",err)
		return err
	}
	
	return nil
}

//查看资产交易记录
func LookRecode(tokenid int64)(string,error){
	//建立连接
	cli,err :=ethclient.Dial(configs.Config.Eth.Connstr)
	if err!= nil{
		fmt.Println("faild to geth",err)
		return "",err
	}
	//函数入口
	instance,err :=NewEths(common.HexToAddress(configs.Config.Eth.QaAddr), cli)
	if err!= nil{
		fmt.Println("faild to NewCallabi",err)
		return "",err
	}
	var ass string
	ass,err =instance.GetAssRecode(nil,big.NewInt(tokenid))
	if err !=nil{
		fmt.Println("faild to GetAssRecode",err)
		return "",err
	}
	return ass,nil
}

//查看版权交易记录
func LookCopyRecode(tokenid int64)(string,error){
	//建立连接
	cli,err :=ethclient.Dial(configs.Config.Eth.Connstr)
	if err!= nil{
		fmt.Println("faild to geth",err)
		return "",err
	}
	//函数入口
	instance,err :=NewEths(common.HexToAddress(configs.Config.Eth.QaAddr), cli)
	if err!= nil{
		fmt.Println("faild to NewCallabi",err)
		return "",err
	}
	var ass string
	ass,err =instance.GetCopyrecode(nil,big.NewInt(tokenid))
	if err !=nil{
		fmt.Println("faild to GetCopyrecode",err)
		return "",err
	}
	return ass,nil
}

//购买信息以太坊记录
func RecodeAss(from,pass,buyer string ,tokenid int64)error{
	//建立连接
	cli,err :=ethclient.Dial(configs.Config.Eth.Connstr)
	if err!= nil{
		fmt.Println("faild to geth",err)
		return err
	}
	//函数入口
	instance,err :=NewEths(common.HexToAddress(configs.Config.Eth.QaAddr), cli)
	if err!= nil{
		fmt.Println("faild to NewCallabi",err)
		return err
	}
	//获取keystore文件名
	fileName,err := utils.GetFileName(string([]rune(from)[2:]),configs.Config.Eth.Keydir)
	if err !=nil{
		fmt.Println("faild to GetFileName",err)
		return err
	}
	//设置签名
	file,err :=os.Open(configs.Config.Eth.Keydir+"/"+fileName)
	if err!= nil{
		fmt.Println("faild to Open",err)
		return err
	}
	
	auth,err :=bind.NewTransactor(file,pass)
	if err!= nil{
		fmt.Println("faild to bind",err)
		return err
	}
	//获取现在时间
	timeStr:=time.Now().Format("2006-01-02 15:04:05")  
	//方法调用
	_,err = instance.AssRecode(auth,buyer+"["+timeStr+"];",big.NewInt(tokenid))
		if err!= nil{
		fmt.Println("faild to AssRecode",err)
		return err
	}
	return nil
}
//版权交易
func CopyTrans(from,pass,toto string ,tokenid int64)error{
	//建立连接
	cli,err :=ethclient.Dial(configs.Config.Eth.Connstr)
	if err!= nil{
		fmt.Println("faild to geth",err)
		return err
	}
	//函数入口
	instance,err :=NewEths(common.HexToAddress(configs.Config.Eth.QaAddr), cli)
	if err!= nil{
		fmt.Println("faild to NewCallabi",err)
		return err
	}
	//获取keystore文件名
	fileName,err := utils.GetFileName(string([]rune(from)[2:]),configs.Config.Eth.Keydir)
	if err !=nil{
		fmt.Println("faild to GetFileName",err)
		return err
	}
	//设置签名
	file,err :=os.Open(configs.Config.Eth.Keydir+"/"+fileName)
	if err!= nil{
		fmt.Println("faild to Open",err)
		return err
	}
	
	auth,err :=bind.NewTransactor(file,pass)
	if err!= nil{
		fmt.Println("faild to bind",err)
		return err
	}
	//方法调用
	_to :=common.HexToAddress(toto)
	_,err =instance.TransferFrom(auth,common.HexToAddress(from),_to,big.NewInt(tokenid))
	if err!= nil{
		fmt.Println("faild to TransferFrom",err)
		return err
	}
	return err
}
//版权交易信息记录
func CoRecode(from,pass,fromname,toname string ,tokenid int64)error{
	//建立连接
	cli,err :=ethclient.Dial(configs.Config.Eth.Connstr)
	if err!= nil{
		fmt.Println("faild to geth",err)
		return err
	}
	//函数入口
	instance,err :=NewEths(common.HexToAddress(configs.Config.Eth.QaAddr), cli)
	if err!= nil{
		fmt.Println("faild to NewCallabi",err)
		return err
	}
	//获取keystore文件名
	fileName,err := utils.GetFileName(string([]rune(from)[2:]),configs.Config.Eth.Keydir)
	if err !=nil{
		fmt.Println("faild to GetFileName",err)
		return err
	}
	//设置签名
	file,err :=os.Open(configs.Config.Eth.Keydir+"/"+fileName)
	if err!= nil{
		fmt.Println("faild to Open",err)
		return err
	}
	
	auth,err :=bind.NewTransactor(file,pass)
	if err!= nil{
		fmt.Println("faild to bind",err)
		return err
	}
	//获取现在时间
	timeStr:=time.Now().Format("2006-01-02 15:04:05")     //获取当前时间，类型是Go的时间类型Time
	//方法调用
	
	_,err = instance.CopyRecode(auth,"["+fromname+"]"+" to "+"["+toname+"]"+" "+timeStr+";",big.NewInt(tokenid))
		if err!= nil{
		fmt.Println("faild to CopyRecode",err)
		return err
	}
	return nil
}

//获取以太币
func MyMoney(myaddress string)(uint64,error){
	//连接geth
	client,err :=rpc.Dial(configs.Config.Eth.Connstr)
	if err !=nil{
		fmt.Println("faild to con to geth",err)
		return 0,err
	}
	//获取以太币
	var money string
	err =client.Call(&money,"eth_getBalance",common.HexToAddress(myaddress),"latest")
	if err !=nil{
		fmt.Println("faild to create eth_getBalance",err)
		return 0,err
	}	
	moneyten :=money[2:]
	n, err := strconv.ParseUint(moneyten, 16, 64)
	n2 := uint64(n)
	return n2,nil
}

//以太币转账
func TranE(from,toto,frompass string)(error){
	//连接geth
	client,err :=rpc.Dial(configs.Config.Eth.Connstr)
	if err !=nil{
		fmt.Println("faild to con to geth",err)
		return err
	}
	//解锁from账户
	var money bool
	err =client.Call(&money,"personal_unlockAccount",common.HexToAddress(from),frompass)
	if err !=nil{
		fmt.Println("失败 to personal_unlockAccount",err)
		return err
	}	
	cli := ethrpc.New("http://localhost:8545")
	_, err = cli.Web3ClientVersion()
	if err !=nil{
		fmt.Println("失败 to web3",err)
		return err
	}
	//转账1eth
	_,errr :=cli.EthSendTransaction(ethrpc.T{ From: from, To: toto, Value: ethrpc.Eth1(), })
	if errr !=nil{
		fmt.Println("失败 to tran",err)
		return errr
	}
	return nil
	
}
