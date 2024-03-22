package main

import (
	"fabric-go-sdk/sdkInit"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

const (
	cc_name    = "simplecc"
	cc_version = "1.0.0"
)

var App sdkInit.Application

func main() {
	/*
		设置默认组织与通道，确定MSP位置，同时设置组织默认值
	*/
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/Org1MSPanchors.tx",
		},
	}
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    "/root/go/src/fabric-go-sdk/chaincode/",
		ChaincodeVersion: cc_version,
	}
	/*
		启动SDK
	*/
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}
	/*
		创建通道，并将节点加入
	*/
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}

	fmt.Println(">> 通过链码外部服务设置链码状态......")
	/*
		初始化链码服务
	*/
	if err := info.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk); err != nil {

		fmt.Println("InitService successful")
		os.Exit(-1)
	}
	/*
		创建代码与区块链服务接口实体
	*/
	App = sdkInit.Application{
		SdkEnvInfo: &info,
	}
	fmt.Println(">> 设置链码状态完成")

	defer info.EvClient.Unregister(sdkInit.BlockListener(info.EvClient))
	defer info.EvClient.Unregister(sdkInit.ChainCodeEventListener(info.EvClient, info.ChaincodeID))
	time.Sleep(time.Second * 10)
	/*
		创建并设置代码与外部的服务接口
	*/
	r := gin.Default()
	r.GET("/userRegister", func(c *gin.Context) {
		acc := c.Query("Account")
		cominfo := c.Query("CompanyInfo")
		bal := c.Query("Balance")
		a := []string{"userRegister", acc, cominfo, bal}
		response, err := App.UserRegister(a)
		fmt.Println(cominfo)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "用户注册失败",
				"data":    err,
				"test1":   acc,
				"test2":   cominfo,
				"test3":   bal,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "用户注册成功，等待管理员审核",
				"data":    response,
			})
		}
	})
	r.GET("/userQuery", func(c *gin.Context) {
		acc := c.Query("Account")
		a := []string{"userQuery", acc}
		response, err := App.UserQuery(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "用户查询失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "用户查询成功",
				"data":    response,
			})
		}
	})
	r.GET("/UserDelete", func(c *gin.Context) {
		acc := c.Query("Account")
		a := []string{"userDelete", acc}
		response, err := App.UserDelete(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "用户删除失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "用户删除成功",
				"data":    response,
			})
		}
	})
	r.GET("/tradeRegister", func(c *gin.Context) {
		id := c.Query("ID")
		from := c.Query("From")
		vol := c.Query("Volume")
		pri := c.Query("Price")
		a := []string{"tradeRegister", id, from, vol, pri}
		response, err := App.TradeRegister(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "交易上传失败",
				"data":    err,
				"test1":   id,
				"test2":   from,
				"test3":   vol,
				"test4":   pri,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "交易上传成功",
				"data":    response,
			})
		}
	})
	r.GET("/tradeQuery", func(c *gin.Context) {
		id := c.Query("ID")
		a := []string{"tradeQuery", id}
		response, err := App.TradeQuery(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "交易查询失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "交易查询成功",
				"data":    response,
			})
		}
	})
	r.GET("/tradeDelete", func(c *gin.Context) {
		id := c.Query("ID")
		a := []string{"tradeDelete", id}
		response, err := App.TradeDelete(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "交易删除失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "交易删除成功",
				"data":    response,
			})
		}
	})
	r.GET("/transaction", func(c *gin.Context) {
		id := c.Query("TradeID")
		to := c.Query("To")
		a := []string{"transaction", id, to}
		response, err := App.Transaction(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "交易失败",
				"data":    err,
				"test1":   id,
				"test2":   to,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "交易达成",
				"data":    response,
			})
		}
	})
	r.GET("/tanReportRegister", func(c *gin.Context) {
		acc := c.Query("Account")
		tanReport := c.Query("TanReport")
		a := []string{"tanReportRegister", acc, tanReport}
		response, err := App.TanReportRegister(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "碳报告上传失败",
				"data":    err,
				"test1":   acc,
				"test2":   tanReport,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "碳报告上传成功",
				"data":    response,
			})
		}
	})

	r.GET("/changeTaoci", func(c *gin.Context) {
		taoci := c.Query("Taoci")
		a := []string{"changeTaoci", taoci}
		response, err := App.ChangeTaoci(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "陶瓷基值更改失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "陶瓷基值更改成功",
				"data":    response,
			})
		}
	})
	r.GET("/changeMg", func(c *gin.Context) {
		mg := c.Query("Mg")
		a := []string{"changeMg", mg}
		response, err := App.ChangeMg(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "镁基值更改失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "镁基值更改成功",
				"data":    response,
			})
		}
	})
	r.GET("/changeED", func(c *gin.Context) {
		ed := c.Query("ED")
		a := []string{"changeED", ed}
		response, err := App.ChangeED(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "碳额度更改失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "碳额度更改成功",
				"data":    response,
			})
		}
	})
	r.GET("/tanHesuan", func(c *gin.Context) {
		acc := c.Query("Account")
		fina := c.Query("Finally")
		huashi1 := c.Query("Huashi1")
		huashi2 := c.Query("Huashi2")
		huashi3 := c.Query("Huashi3")
		mg1 := c.Query("Mg")
		a := []string{"tanHesuan", acc, fina, huashi1, huashi2, huashi3, mg1}
		response, err := App.TanHesuan(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "碳核算失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "碳核算成功",
				"data":    response,
			})
		}
	})
	r.GET("/tanHesuanTXT", func(c *gin.Context) {
		acc := c.Query("Account")
		went := c.Query("Wenshi")
		huot := c.Query("Huodong")
		pait := c.Query("Paifang")
		a := []string{"tanHesuanTXT", acc, went, huot, pait}
		response, err := App.TanHesuanTXT(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "文本写入失败",
				"data":    err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "文本写入成功",
				"data":    response,
			})
		}
	})
	r.GET("/proceedRegister", func(c *gin.Context) {
		pid := c.Query("PrID")
		tid := c.Query("TradeID")
		from := c.Query("From")
		to := c.Query("To")
		a := []string{"proceedRegister", pid, tid, from, to}
		response, err := App.ProceedRegister(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "诉讼上传失败",
				"data":    err,
				"test1":   pid,
				"test2":   tid,
				"test3":   from,
				"test4":   to,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "诉讼上传成功",
				"data":    response,
			})
		}
	})
	r.GET("/proceed", func(c *gin.Context) {
		pid := c.Query("PrID")
		fin := c.Query("Final")
		a := []string{"proceed", pid, fin}
		response, err := App.Proceed(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  "failed",
				"message": "诉讼处理失败",
				"data":    err,
				"test1":   pid,
				"test2":   fin,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "诉讼处理成功",
				"data":    response,
			})
		}
	})

	r.Run(":9090")
}
