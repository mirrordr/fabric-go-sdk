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
	r.GET("/k2RawSignRegister", func(c *gin.Context) {
		id := c.Query("ID")
		k2 := c.Query("K2")
		rawsign := c.Query("Rawsign")
		a := []string{"K2RawSignRegister", id, k2, rawsign}
		response, err := App.K2RawSignRegister(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": response,
				"final":  "success",
			})
		}
	})
	r.GET("/k2RawSignQuery", func(c *gin.Context) {
		id := c.Query("ID")
		a := []string{"K2RawSignQuery", id}
		response, err := App.K2RawSignQuery(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": response,
				"final":  "success",
			})
		}
	})
	r.GET("/k2RawSignDelete", func(c *gin.Context) {
		id := c.Query("ID")
		a := []string{"K2RawSignDelete", id}
		response, err := App.K2RawSignDelete(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": response,
				"final":  "success",
			})
		}
	})
	r.GET("/companyViewRecordRegister", func(c *gin.Context) {
		id := c.Query("ID")
		resumeID := c.Query("ResumeID")
		schoolCode := c.Query("SchoolCode")
		staffID := c.Query("StaffID")
		companyID := c.Query("CompanyID")
		a := []string{"CompanyViewRecordRegister", id, resumeID, schoolCode, staffID, companyID}
		response, err := App.CompanyViewRecordRegister(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": response,
				"final":  "success",
			})
		}
	})
	r.GET("/companyViewRecordQuery", func(c *gin.Context) {
		id := c.Query("ID")
		a := []string{"CompanyViewRecordQuery", id}
		response, err := App.CompanyViewRecordQuery(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": response,
				"final":  "success",
			})
		}
	})
	r.GET("/companyViewRecordDelete", func(c *gin.Context) {
		id := c.Query("ID")
		a := []string{"CompanyViewRecordDelete", id}
		response, err := App.CompanyViewRecordDelete(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"result": response,
				"final":  "success",
			})
		}
	})
	r.Run(":9090")
}
