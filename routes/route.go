package routes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"go-copyright-p1/configs"
	"go-copyright-p1/dbs"
	"go-copyright-p1/eths"
	"go-copyright-p1/utils"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

const PAGE_MAX_PIC = 10
const BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

func PingHandler(c echo.Context) error {

	return c.String(http.StatusOK, "pong")
}

//注册
func Register(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	account := &dbs.Account{}
	if err := c.Bind(account); err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	//操作geth创建账户
	address, err := eths.NewAcc(account.IdentityID, configs.Config.Eth.Connstr)
	if err != nil {
		resp.Errno = utils.RECODE_IPCERR
		return err
	}
	//账户信息插入数据库
	//insert into account(email,username,identity_id,address) values()
	sql := fmt.Sprintf("insert into account(email,username,identity_id,address) values('%s','%s','%s','%s')",
		account.Email,
		account.UserName,
		account.IdentityID,
		address,
	)
	fmt.Println(sql)
	_, err = dbs.Create(sql)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//session处理
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["name"] = account.UserName
	sess.Values["address"] = address
	sess.Values["password"] = account.IdentityID
	sess.Save(c.Request(), c.Response())
	return nil
}

//session获取
func GetSession(c echo.Context) error {
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"]
	name := sess.Values["name"]
	password := sess.Values["password"]
	if address == nil {
		fmt.Println("faild to get session,address,nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	if name == nil {
		fmt.Println("失败 to get session,address,nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	mymoney, errr := eths.MyMoney(address.(string))
	str1 := strconv.FormatUint(mymoney, 10)
	if errr != nil {
		fmt.Println("失败 to get 以太币")
		resp.Errno = utils.RECODE_GETMONEY
		return errr
	}
	mapResp := make(map[string]interface{})
	mapResp["name"] = name
	mapResp["address"] = address
	mapResp["password"] = password
	mapResp["money"] = str1
	resp.Data = mapResp
	return nil
}

//登录
func Login(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	account := &dbs.Account{}
	if err := c.Bind(account); err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	//查询数据
	//select * from account where username='qly' and identity_id='123456'
	sql := fmt.Sprintf("select * from account where username='%s' and identity_id='%s'",
		account.UserName,
		account.IdentityID,
	)
	fmt.Println(sql)
	m, n, err := dbs.DBQuery(sql)
	if err != nil || n <= 0 {
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	rows := m[0]
	//session处理
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["name"] = account.UserName
	sess.Values["address"] = rows["address"]
	sess.Values["password"] = account.IdentityID
	sess.Save(c.Request(), c.Response())
	return nil
}

//上传资产

func Upload(c echo.Context) error {
	//
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//
	content := &dbs.Content{}
	h, err := c.FormFile("fileName")
	if err != nil {
		fmt.Println("faild to formfile", err)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	src, err := h.Open()
	if err != nil {
		fmt.Println("open file err", err)
		return err
	}
	defer src.Close()
	//计算资产hsah
	cData := make([]byte, h.Size)
	n, err := src.Read(cData)
	if err != nil || h.Size != int64(n) {
		resp.Errno = utils.RECODE_IOERR
		return err
	}
	content.ContentHash = fmt.Sprintf("%x", sha256.Sum256(cData))
	//查重
	//select 1 from conetnt where content_hash='%s'
	sql := fmt.Sprintf("select 1 from content where content_hash='%s'",
		content.ContentHash,
	)
	fmt.Println(sql)
	_, n, errr := dbs.DBQuery(sql)
	if errr != nil || n > 0 {
		resp.Errno = utils.RECODE_DOUBLE
		return errr
	}
	//加密
	//1,生成随机秘钥
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(88888888) + 10000000
	passone := strconv.Itoa(rnd)
	copass := passone + passone + passone
	//2，用秘钥加密
	encryptCode := AesEncrypt(cData, copass)
	//存入服务器
	content.Content = "./static/photo/" + h.Filename
	dst, err := os.Create(content.Content)
	if err != nil {
		fmt.Println("faild to create f", err, content.Content)
		resp.Errno = utils.RECODE_IOERR
		return err
	}
	defer dst.Close()
	dst.Write(encryptCode)
	content.Title = h.Filename
	cocopass := BaseEncode(copass)
	//登记秘钥
	//insert into content_pass(content_hash,cpass) values
	sql = fmt.Sprintf("insert into content_pass(content_hash,cpass) values('%s','%s')",
		content.ContentHash,
		cocopass,
	)
	fmt.Println(sql)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to insert content_pass")
		resp.Errno = utils.RECODE_DBERR
		return err
	}

	//写入资产
	//操作数据库插入
	content.AddContent()

	sess, _ := session.Get("session", c)
	username := sess.Values["name"].(string)
	fromAddr, ok := sess.Values["address"].(string)
	fromPassword := sess.Values["password"].(string)
	if fromAddr == "" || fromPassword == "" || !ok {
		resp.Errno = utils.RECODE_SESSIONERR
		return errors.New("no session")
	}
	//登记私人秘钥，默认123，可在页面修改
	//insert into account_pass(username,content_hash,apass) values
	sql = fmt.Sprintf("insert into account_pass(username,content_hash,apass) values('%s','%s','%s')",
		username,
		content.ContentHash,
		BaseEncode("123"),
	)
	fmt.Println(sql)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to insert account_pass")
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//操作以太坊
	go eths.UploadPic(fromAddr, fromPassword, content.ContentHash, content.Title)

	return nil
}

//查看当前用户资产

func GetContents(c echo.Context) error {
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"]
	if address == nil {
		fmt.Println("faild to get session,address,nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	//select a.content_hash,title,token_id from content a,account_content b where a.content_hash=b.content_hash and address=''
	sql := fmt.Sprintf("select a.content_hash,title,token_id from content a,account_content b where a.content_hash=b.content_hash and address='%s'", address)
	fmt.Println(sql)
	contents, num, err := dbs.DBQuery(sql)
	if err != nil || num <= 0 {
		resp.Errno = utils.RECODE_DBERR
		fmt.Println("faild to query my zichan", err)
		return err
	}
	total_page := int(num)/PAGE_MAX_PIC + 1
	current_page := 1
	mapResp := make(map[string]interface{})
	mapResp["total_page"] = total_page
	mapResp["current_page"] = current_page
	mapResp["contents"] = contents

	resp.Data = mapResp
	return nil
}

//发起交易

func Auction(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	auction := &dbs.Auction{}
	if err := c.Bind(auction); err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	//查询是否重复
	//select 1 from auction where token_id='%d'
	sql := fmt.Sprintf("select 1 from auction where token_id='%d'", auction.TokenID)
	fmt.Println(sql)
	_, n, _ := dbs.DBQuery(sql)
	if n == 0 {
		//插入数据
		//insert into auction(content_hash,address,status) values
		sql = fmt.Sprintf("insert into auction(content_hash,token_id,status) values('%s','%d',1)",
			auction.ContentHash,
			auction.TokenID,
		)
		fmt.Println(sql)
		_, err := dbs.Create(sql)
		if err != nil {
			fmt.Println("faild to insert auction")
			resp.Errno = utils.RECODE_DBERR
			return err
		}

	}

	return nil
}

//查看资产交易记录
func Tranrecode(c echo.Context) error {
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)
	//解析数据
	auction := &dbs.Auction{}
	if err := c.Bind(auction); err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	//操作以太坊

	ass, err := eths.LookRecode(auction.TokenID)
	if err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		fmt.Println("faild to go lookRecode")
		return err
	}

	mapResp := make(map[string]interface{})
	mapResp["ass"] = ass
	resp.Data = mapResp

	return nil
}

//查看版权交易记录
func CopyTranrecode(c echo.Context) error {
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)
	//解析数据
	auction := &dbs.Auction{}
	if err := c.Bind(auction); err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}
	//操作以太坊

	cop, err := eths.LookCopyRecode(auction.TokenID)
	if err != nil {
		resp.Errno = utils.RECODE_PARAMERR
		fmt.Println("faild to go LookCopyRecode")
		return err
	}

	mapResp := make(map[string]interface{})
	mapResp["cop"] = cop
	resp.Data = mapResp

	return nil
}

//查看交易

func GetAuction(c echo.Context) error {
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)
	owner := c.QueryParam("own")
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"]
	if address == nil {
		fmt.Println("faild to get session,address,nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	//select username,title,au.content_hash,au.token_id from account a,account_content aa,content c,auction au where au.content_hash=c.content_hash and c.content_hash=aa.content_hash and aa.address=a.address and au.`status`=1 and username='qly'
	var sql string
	if owner == "" {
		sql = fmt.Sprintf("select username,title,au.content_hash,au.token_id from account a,account_content aa,content c,auction au where au.content_hash=c.content_hash and c.content_hash=aa.content_hash and aa.address=a.address and au.`status`=1")
	} else {
		sql = fmt.Sprintf("select username,title,au.content_hash,au.token_id from account a,account_content aa,content c,auction au where au.content_hash=c.content_hash and c.content_hash=aa.content_hash and aa.address=a.address and au.`status`=1 and username='%s'", owner)
	}
	fmt.Println(sql)
	auctions, _, err := dbs.DBQuery(sql)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		fmt.Println("faild to query my paimai", err)
		return err
	}

	resp.Data = auctions
	return nil
}

//购买资产
func Buy(c echo.Context) error {
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)
	//解析数据
	tokenid := c.QueryParam("tokenid")
	hhash := c.QueryParam("hhash")
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"].(string)
	name := sess.Values["name"].(string)
	pass := sess.Values["password"].(string)
	if address == "" {
		fmt.Println("faild to get session,address,nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	//查找秘钥
	//select cpass,apass,username from content_pass c,account_pass a where a.content_hash=c.content_hash and c.content_hash='%s'
	sql := fmt.Sprintf("select cpass,apass,username from content_pass c,account_pass a where a.content_hash=c.content_hash and c.content_hash='%s'", hhash)
	passlist, _, errr := dbs.DBQuery(sql)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		fmt.Println("faild to query my daiqueren", errr)
		return err
	}
	lookpass := passlist[0]
	mapResp := make(map[string]interface{})
	mapResp["cpass"] = BaseDecode(lookpass["cpass"])
	mapResp["apass"] = BaseDecode(lookpass["apass"])

	resp.Data = mapResp

	//操作以太坊
	go func() {
		//以太坊记录购买信息
		_tokenid, _ := strconv.ParseInt(tokenid, 10, 32)
		err = eths.RecodeAss(address, pass, name, _tokenid)
		if err != nil {
			fmt.Println("faild to go recodeAss", err)
			return
		}
	}()
	if err != nil {
		return err
	}

	return nil
}

//提交申请
func Apply(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	tokenid := c.QueryParam("tokenid")
	hash := c.QueryParam("hash")
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"].(string)
	if address == "" {
		fmt.Println("faild to get session,address,nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	//数据库操作
	//insert into assetok(content_hash,token_id,address) values()
	sql := fmt.Sprintf("insert into assetok(content_hash,token_id,address) values('%s','%s','%s')",
		hash,
		tokenid,
		address,
	)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to insert assetok")
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//update auction set `status`=2 where token_id='%s'
	sql = fmt.Sprintf("update auction set `status`=2 where token_id='%s'", tokenid)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to update auction")
		resp.Errno = utils.RECODE_DBERR
		return err
	}

	return nil
}

//取消交易
func Revoke(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	tokenid := c.QueryParam("tokenid")
	//数据库操作
	//delete from auction where token_id='%s'
	sql := fmt.Sprintf("delete from auction where token_id='%s'", tokenid)
	_, err := dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to delete auction")
		resp.Errno = utils.RECODE_DBERR
		return err
	}

	return nil
}

//查看待确认交易
func OkList(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"].(string)
	if address == "" {
		fmt.Println("faild to get session,address,nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	//数据库操作
	//SELECT a.address,a.username,e.title,c.token_id from account a,assetok c,auction d,content e where a.address=c.address and c.content_hash=e.content_hash and d.token_id=c.token_id and d.`status`=2 and c.token_id in(select b.token_id from account_content b where b.address='%s')
	sql := fmt.Sprintf("SELECT a.address,a.username,e.title,c.token_id from account a,assetok c,auction d,content e where a.address=c.address and c.content_hash=e.content_hash and d.token_id=c.token_id and d.`status`=2 and c.token_id in(select b.token_id from account_content b where b.address='%s')", address)
	okli, _, err := dbs.DBQuery(sql)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		fmt.Println("faild to query my daiqueren", err)
		return err
	}

	resp.Data = okli
	return nil
}

//拒绝交易版权
func Notrans(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	tokenid := c.QueryParam("tokenid")
	//数据库操作
	//delete from assetok where token_id='%s'
	sql := fmt.Sprintf("delete from assetok where token_id='%s'", tokenid)
	_, err := dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to delete assetok")
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//update auction set `status`=1 where token_id='%s'
	sql = fmt.Sprintf("update auction set `status`=1 where token_id='%s'", tokenid)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to update auction")
		resp.Errno = utils.RECODE_DBERR
		return err
	}

	return nil
}

//确认交易版权
func ConfirmTran(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	tokenid := c.QueryParam("tokenid")
	address1 := c.QueryParam("addr")
	okname := c.QueryParam("okname")
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"].(string)
	name := sess.Values["name"].(string)
	pass := sess.Values["password"].(string)
	if address == "" {
		fmt.Println("faild to get session,address,nil")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	//select * from account where username='%s'
	sql := fmt.Sprintf("select * from account where username='%s'", okname)
	m, _, errr := dbs.DBQuery(sql)
	if errr != nil {
		resp.Errno = utils.RECODE_DBERR
		fmt.Println("faild to query my content_pass", errr)
		return errr
	}
	rows := m[0]
	topass := rows["identity_id"]
	//以太坊操作
	_tokenid, _ := strconv.ParseInt(tokenid, 10, 32)
	go func() {
		eths.CopyTrans(address, pass, address1, _tokenid)
		if err != nil {
			fmt.Println("faild to go CopyTrans", err)
			return
		}
		eths.TranE(address1, address, topass)
		if err != nil {
			fmt.Println("faild to go TranE", err)
			return
		}
		eths.CoRecode(address, pass, name, okname, _tokenid)
		if err != nil {
			fmt.Println("faild to go CoRecode", err)
			return
		}
	}()

	//数据库操作
	//delete from assetok where token_id='%s'
	sql = fmt.Sprintf("delete from assetok where token_id='%s'", tokenid)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to delete assetok")
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//delete from auction where token_id='%s'
	sql = fmt.Sprintf("delete from auction where token_id='%s'", tokenid)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to delete auction")
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	//update account_content set address='%s' where token_id='%s'
	sql = fmt.Sprintf("update account_content set address='%s' where token_id='%s'", address1, tokenid)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to update account_content")
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	return nil
}

//写文件到服务器本地
func Download(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)

	//解析数据
	title := c.QueryParam("Title")
	cpass := c.QueryParam("cpass")
	apass := c.QueryParam("apass")
	//判断秘钥
	//select 1 from content_pass c,account_pass a where a.content_hash=c.content_hash and cpass='%s' and apass='%s'
	sql := fmt.Sprintf("select 1 from content_pass c,account_pass a where a.content_hash=c.content_hash and cpass='%s' and apass='%s'", BaseEncode(cpass), BaseEncode(apass))
	_, n, err := dbs.DBQuery(sql)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		fmt.Println("faild to query my content_pass", err)
		return err
	}
	if n <= 0 {
		return err
	}

	//写文件到服务器本地

	src, err := os.Open("static/photo/" + title)
	if err != nil {
		resp.Errno = utils.RECODE_IOERR
		fmt.Println("FFF to open f", err)
		return err
	}
	defer src.Close()

	var result int64
	err = filepath.Walk("static/photo/"+title, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	cData := make([]byte, result)
	_, err = src.Read(cData)
	if err != nil || cData == nil {
		fmt.Println("FFF to read f", err)
		resp.Errno = utils.RECODE_IOERR
		return err
	}
	//解密
	decryptCode := AesDecrypt(cData, cpass)

	dst, err := os.Create("c:\\asset\\" + title)
	if err != nil {
		resp.Errno = utils.RECODE_IOERR
		fmt.Println("FFF to create f", err)
		return err
	}
	defer dst.Close()
	dst.Write(decryptCode)
	if err != nil {
		resp.Errno = utils.RECODE_IOERR
		fmt.Println("FFF to wirte f", err)
		return err
	}
	return nil
}

//下载文件
func DownloadFile(c echo.Context) error {
	title := c.QueryParam("Title")
	defer os.Remove("c:\\asset\\" + title)
	return c.File("c:\\asset\\" + title)
}

//设置秘钥
func SetPass(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	ahash := c.QueryParam("hhash")
	aapass := c.QueryParam("aapass")
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	name := sess.Values["name"].(string)
	aaapass := BaseEncode(aapass)
	//update account_pass set apass='%s',username='%s' where content_hash='%s'
	sql := fmt.Sprintf("update account_pass set apass='%s',username='%s' where content_hash='%s'", aaapass, name, ahash)
	_, err = dbs.Create(sql)
	if err != nil {
		fmt.Println("faild to update account_pass")
		resp.Errno = utils.RECODE_DBERR
		return err
	}
	return nil
}

//查看秘钥
func GetPass(c echo.Context) error {
	//响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK

	defer utils.ResponseData(c, &resp)
	//解析数据
	ahash := c.QueryParam("hhash")
	//处理session
	sess, err := session.Get("session", c)
	if err != nil {
		fmt.Println("faild to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	name := sess.Values["name"].(string)
	//select cpass,apass,username from content_pass c,account_pass a where a.content_hash=c.content_hash and c.content_hash='%s'
	sql := fmt.Sprintf("select cpass,apass,username from content_pass c,account_pass a where a.content_hash=c.content_hash and c.content_hash='%s'", ahash)
	passlist, _, errr := dbs.DBQuery(sql)
	if errr != nil {
		resp.Errno = utils.RECODE_DBERR
		fmt.Println("faild to query my daiqueren", errr)
		return errr
	}
	lookpass := passlist[0]
	mapResp := make(map[string]interface{})
	if name != lookpass["username"] {
		mapResp["okok"] = "请立即修改私人密码"
	} else {
		mapResp["okok"] = ""
	}
	mapResp["cpass"] = BaseDecode(lookpass["cpass"])
	mapResp["apass"] = BaseDecode(lookpass["apass"])

	resp.Data = mapResp
	return nil
}

//以下全为AES加密算法
func AesEncrypt(origData []byte, key string) []byte {
	// 转成字节数组
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	return cryted

}

//AES解密
func AesDecrypt(crytedByte []byte, key string) []byte {
	// 转成字节数组
	k := []byte(key)

	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return orig
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//Base64加密
func BaseEncode(data string) string {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&data))))
	coder := base64.NewEncoding(BASE64Table)
	return coder.EncodeToString(content)
}

//Base64解密
func BaseDecode(data string) string {
	coder := base64.NewEncoding(BASE64Table)
	result, _ := coder.DecodeString(data)
	return *(*string)(unsafe.Pointer(&result))
}
