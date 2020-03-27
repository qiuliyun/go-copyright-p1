//监控事件
package eths

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-copyright-p1/dbs"
	"reflect"
	"strconv"
)

func LogDataUnpack(start, end int, val interface{}, data []byte) (err error) {
	length := len(data)
	fmt.Println("call--- LogDataUnpack begin", reflect.TypeOf(val).String(), length)

	if start >= length || end > length {
		return errors.New("too short datas")
	}
	pdata := data[start:end]
	

	fmt.Println(string(data), string(pdata))
	if reflect.TypeOf(val).String() == "int64" || reflect.TypeOf(val).String() == "*int64" {
		var tmpval *int64 = val.(*int64)
		*tmpval, err = strconv.ParseInt(string(pdata), 16, 32)
		fmt.Println("call ParseInt", val)
	} else if reflect.TypeOf(val).String() == "string" || reflect.TypeOf(val).String() == "*string" {
		var tmpval *string = val.(*string)
		*tmpval = string(pdata)
		fmt.Println("call ParseInt", val)
	}

	fmt.Println("call--- LogDataUnpack end", val)
	return nil
}
func ParseMintEvent2Db(data []byte) error {
	fmt.Println(string(data))
	var tokenId int64
	err := LogDataUnpack(32*5, 32*6, &tokenId, data)
	if err != nil {
		fmt.Println("faile to get tokenid", err)
		return err
	}
	fmt.Println("tokenid===", tokenId)
	var pixHash string
	err = LogDataUnpack(32*0, 32*2, &pixHash, data)
	if err != nil {
		return err
	}
	fmt.Println("pixHash===", pixHash)
	var pixAddr string
	err = LogDataUnpack(88, 128, &pixAddr, data)
	if err != nil {
		return err
	}
	pixAddr = "0x" + pixAddr
	fmt.Println("pixAddr===", pixAddr)
	//插入到数据库中
	sql := fmt.Sprintf("insert into account_content(content_hash,token_id,address) values('%s',%d,'%s')", pixHash, tokenId, pixAddr)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("failed to insert into mysql ", sql, err)
		return err
	}
	return nil
}

func EventSubscrib(connstr,contractAddr string) error{
	cli,err :=ethclient.Dial(connstr)
	if err !=nil{
		fmt.Println("faild to geth",err)
		return err
	}
	//合约地址处理
	cAddress :=common.HexToAddress(contractAddr)
	newAssetHash :=crypto.Keccak256Hash([]byte("newAsset(bytes32,address,uint256)"))
	//过滤处理
	query :=ethereum.FilterQuery{
		Addresses : []common.Address{cAddress},
		Topics : [][]common.Hash{{newAssetHash}},
	}
	//通道
	logs :=make(chan types.Log)
	//订阅
	sub,err :=cli.SubscribeFilterLogs(context.Background(),query,logs)
		if err !=nil{
		fmt.Println("faild to SubscribeFilterLogs",err)
		return err
	}
	//订阅返回处理
	for{
		select {
			case err :=<-sub.Err():
				fmt.Println("get sub err",err)
			case vLog :=<-logs:
				data,err :=vLog.MarshalJSON()
				fmt.Println(string(data),err)
				ParseMintEvent2Db([]byte(common.Bytes2Hex(vLog.Data)))
		}
	}
}
