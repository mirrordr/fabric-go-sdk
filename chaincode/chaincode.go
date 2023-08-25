package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type SimpleAsset struct {
}

/*
保存用户和对应的k2
ID对应用户的区块链id
K2为对应用户的密钥一部分
RawSign为简历的明文签名
*/
type k2RawSign struct {
	ID      string `json:"ID"`
	K2      string `json:"K2"`
	RawSign []byte `json:"RawSign"`
}

/*
用户简历对应信息被访问的信息保存
Nomber为被访问次数
IdReadRecordsMap为查询信息
*/
type companyViewRecordMap struct {
	ID               string                    `json:"ID"`
	Number           int                       `json:"Number"`
	IdReadRecordsMap map[int]companyViewRecord `json:"IdReadRecordsMap"`
}

/*
ResumeID为简历id
ReadAt为简历被查询时间
SchoolCode为学校id
StaffID为学号
CompanyID为公司ID
*/
type companyViewRecord struct {
	ResumeID   string    `json:"ResumeID"`
	ReadAt     time.Time `json:"ReadAt"`
	SchoolCode string    `json:"SchoolCode"`
	StaffID    string    `json:"StaffID"`
	CompanyID  string    `json:"CompanyID"`
}

// Init /*区块链的初始化
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Printf("init...")
	return shim.Success(nil)
}

// Invoke /*调用区块链的函数
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	switch fn {
	case "k2RawSignRegister":
		return t.k2RawSignRegister(stub, args)
	case "k2RawSignQuery":
		return t.k2RawSignQuery(stub, args)
	case "k2RawSignDelete":
		return t.k2RawSignDelete(stub, args)
	case "companyViewRecordRegister":
		return t.companyViewRecordRegister(stub, args)
	case "companyViewRecordQuery":
		return t.companyViewRecordQuery(stub, args)
	case "companyViewRecordDelete":
		return t.companyViewRecordDelete(stub, args)
	default:
		return shim.Error("Unsupported function")
	}
	return shim.Success(nil)
}

/*
创建一个id对应的k2存储结构体并实现上链
*/
func (t *SimpleAsset) k2RawSignRegister(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of args.Expecting 3")
	}
	id := args[0]
	k2 := args[1]
	rawsign := args[2]
	if id == "" || k2 == "" || rawsign == "" {
		return shim.Error("Invalid args.")
	}
	idByes, err := stub.GetState(id)
	if err == nil && len(idByes) != 0 {
		return shim.Error("id already exists")
	}
	rawSign := []byte(rawsign)
	k2rawsign := k2RawSign{
		ID:      id,
		K2:      k2,
		RawSign: rawSign,
	}
	userByes, err := json.Marshal(k2rawsign)
	if err != nil {
		return shim.Error("marshal k2rawsign error")
	}
	if err = stub.PutState(id, userByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

/*
输入id从区块链中取出对应的k2结构体
*/
func (t *SimpleAsset) k2RawSignQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	id := args[0]
	idBytes, err := stub.GetState(id)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	return shim.Success(idBytes)
}

/*
从区块链中删除id对应的结构体
*/
func (t *SimpleAsset) k2RawSignDelete(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	id := args[0]
	if err := stub.DelState(id); err != nil {
		return shim.Error("Failed to delete k2rawsign")
	}
	return shim.Success(nil)
}

/*
通过输入相关数据，同时输入的用户id要在最后添加3个0
会查询是否存在记录，不存在重新创建公司查看履历记录表
若存在则在原记录后添加新记录
*/
func (t *SimpleAsset) companyViewRecordRegister(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of args.Expecting 5")
	}
	id := args[0]
	resumeID := args[1]
	schoolCode := args[2]
	staffID := args[3]
	companyID := args[4]
	if id == "" || resumeID == "" || schoolCode == "" || staffID == "" || companyID == "" {
		return shim.Error("Invalid args.")
	}
	idByes, err := stub.GetState(id)
	if err != nil {
		return shim.Error("Search error!!")
	}
	var recordMap companyViewRecordMap
	if len(idByes) != 0 {
		time1 := time.Now()
		if err = json.Unmarshal(idByes, &recordMap); err != nil {
			return shim.Error("Failed to unmarshal recordMap")
		}
		record := companyViewRecord{
			ResumeID:   resumeID,
			ReadAt:     time1,
			SchoolCode: schoolCode,
			StaffID:    staffID,
			CompanyID:  companyID,
		}
		recordMap.ID = id
		recordMap.Number = recordMap.Number + 1
		recordMap.IdReadRecordsMap[recordMap.Number] = record
		newrecordMap, err := json.Marshal(recordMap)
		if err != nil {
			return shim.Error("marshal recordMap error")
		}
		if err = stub.PutState(id, newrecordMap); err != nil {
			return shim.Error("Failed to put state")
		}
		return shim.Success(nil)
	} else {
		recordMap.IdReadRecordsMap = make(map[int]companyViewRecord)
		time1 := time.Now()
		record := companyViewRecord{
			ResumeID:   resumeID,
			ReadAt:     time1,
			SchoolCode: schoolCode,
			StaffID:    staffID,
			CompanyID:  companyID,
		}
		recordMap.ID = id
		recordMap.Number = 0
		recordMap.IdReadRecordsMap[0] = record
		newrecordMap, err := json.Marshal(recordMap)
		if err != nil {
			return shim.Error("marshal recordMap error")
		}
		if err = stub.PutState(id, newrecordMap); err != nil {
			return shim.Error("Failed to put state")
		}
		return shim.Success(nil)
	}

}

/*
查询用户简历查看记录，用户id同样加3个0
*/
func (t *SimpleAsset) companyViewRecordQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	id := args[0]
	idBytes, err := stub.GetState(id)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	return shim.Success(idBytes)
}

/*
删除对应用户id的简历查看记录
*/
func (t *SimpleAsset) companyViewRecordDelete(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	id := args[0]
	if err := stub.DelState(id); err != nil {
		return shim.Error("Failed to delete companyViewRecordMap")
	}
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
