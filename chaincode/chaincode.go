package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
)

type SimpleAsset struct {
}

/*
保存用户
*/
type User struct {
	Account     string      `json:"Account"`
	CompanyInfo CompanyInfo `json:"CompanyInfo"`
	Balance     int64       `json:"Balance"`
}

/*
公司基本信息
*/
type CompanyInfo struct {
	Name  string `json:"Name"`
	Type  string `json:"Type"`
	Owner string `json:"Owner"`
}

/*
订单信息
*/
type Trade struct {
	TradeId     string `json:"TradeId"`
	FromAccount string `json:"FromAccount"`
	Volume      int64  `json:"Volume"`
	Price       int64  `json:"Price"`
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
	case "userRegister":
		return t.UserRegister(stub, args)
	case "userQuery":
		return t.UserQuery(stub, args)
	case "userDelete":
		return t.UserDelete(stub, args)
	case "tradeRegister":
		return t.TradeRegister(stub, args)
	case "tradeQuery":
		return t.TradeQuery(stub, args)
	case "tradeDelete":
		return t.TradeDelete(stub, args)
	case "transaction":
		return t.Transaction(stub, args)

	default:
		return shim.Error("Unsupported function")
	}
	return shim.Success(nil)
}

/*
创建一个account对应的User存储结构体并实现上链
*/
func (t *SimpleAsset) UserRegister(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of args.Expecting 5")
	}
	acc := args[0]
	bal := args[1]
	nam := args[2]
	typ := args[3]
	own := args[4]
	if acc == "" || bal == "" || nam == "" || typ == "" || own == "" {
		return shim.Error("Invalid args.")
	}
	accountByes, err := stub.GetState(acc)
	if err == nil && len(accountByes) != 0 {
		return shim.Error("account already exists")
	}
	balance, _ := strconv.ParseInt(bal, 10, 64)
	company := CompanyInfo{
		Name:  nam,
		Type:  typ,
		Owner: own,
	}
	user := User{
		Account:     acc,
		CompanyInfo: company,
		Balance:     balance,
	}
	userByes, err := json.Marshal(user)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(acc, userByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

/*
输入account从区块链中取出对应的User结构体
*/
func (t *SimpleAsset) UserQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	acc := args[0]
	idBytes, err := stub.GetState(acc)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	return shim.Success(idBytes)
}

/*
从区块链中删除account对应的结构体
*/
func (t *SimpleAsset) UserDelete(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	acc := args[0]
	if err := stub.DelState(acc); err != nil {
		return shim.Error("Failed to delete k2rawsign")
	}
	return shim.Success(nil)
}

/*
创建一个tradeid对应的Trade存储结构体并实现上链
*/
func (t *SimpleAsset) TradeRegister(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of args.Expecting 4")
	}
	id := args[0]
	from := args[1]
	vol := args[2]
	pri := args[3]
	if id == "" || from == "" || vol == "" || pri == "" {
		return shim.Error("Invalid args.")
	}
	tradeByes, err := stub.GetState(id)
	if err != nil && len(tradeByes) != 0 {
		return shim.Error("Search error!!")
	}
	volume, _ := strconv.ParseInt(vol, 10, 64)
	price, _ := strconv.ParseInt(pri, 10, 64)
	trade := Trade{
		TradeId:     id,
		FromAccount: from,
		Volume:      volume,
		Price:       price,
	}
	traByes, err := json.Marshal(trade)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(id, traByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)

}

/*
查询交易信息
*/
func (t *SimpleAsset) TradeQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
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
删除交易信息
*/

func (t *SimpleAsset) TradeDelete(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	id := args[0]
	if err := stub.DelState(id); err != nil {
		return shim.Error("Failed to delete companyViewRecordMap")
	}
	return shim.Success(nil)
}

/*
执行交易from->to
*/
func (t *SimpleAsset) Transaction(stub shim.ChaincodeStubInterface, args []string) peer.Response { //执行资金的密态转移
	if len(args) != 3 {
		return shim.Error("Incorrect number of args.Expecting 3")
	}
	userFrom, userTo, bal := args[0], args[1], args[2]
	existFrom, err := stub.GetState(userFrom)
	if err == nil && len(existFrom) == 0 {
		return shim.Error("sender does not exist")
	}
	existTo, err := stub.GetState(userTo)
	if err == nil && len(existTo) == 0 {
		return shim.Error("receiver does not exist")
	}
	var from, to User
	if userFrom == "" || userTo == "" || bal == "" {
		return shim.Error("Invalid args")
	}
	balance, _ := strconv.ParseInt(bal, 10, 64)
	userFromBytes, err := stub.GetState(userFrom)
	if err != nil {
		return shim.Error("Failed to get userFrom state")
	}
	if err = json.Unmarshal(userFromBytes, &from); err != nil {
		return shim.Error("Failed to unmarshal userFrom")
	}
	userToByes, err := stub.GetState(userTo)
	if err != nil {
		return shim.Error("Failed to get userFrom state")
	}
	if err = json.Unmarshal(userToByes, &to); err != nil {
		return shim.Error("Failed to unmarshal userFrom")
	}
	from.Balance = from.Balance - balance
	to.Balance = to.Balance + balance
	newFrom, err := json.Marshal(from)
	if err != nil {
		return shim.Error("marshal user error")
	}
	newTo, err := json.Marshal(to)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(from.Account, newFrom); err != nil {
		return shim.Error("Failed to put state")
	}
	if err = stub.PutState(to.Account, newTo); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
