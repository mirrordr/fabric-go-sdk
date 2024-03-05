package main

import (
	"chaincode/tanhesuan"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
	"time"
)

type SimpleAsset struct {
}

type TanReport struct {
	Huashi        tanhesuan.Fossil_Fuel_Combustion                         `json:"Huashi"`
	Taocizhuanyou tanhesuan.Ceramics_Indusry_Production_Process            `json:"Taocizhuanyou"`
	Dianli        tanhesuan.Electricity_And_Heat_Emissions                 `json:"Dianli"`
	Ma            tanhesuan.Magnesium_smelting_Industry_Production_Process `json:"Ma"`
	Time          time.Time                                                `json:"Time"`
	Examine       Examine                                                  `json:"Examine"`
}

/*
保存用户
*/
type User struct {
	Account     string              `json:"Account"`
	CompanyInfo CompanyInfo         `json:"CompanyInfo"`
	Balance     float64             `json:"Balance"`
	Examine     Examine             `json:"Examine"`
	FromNumber  int64               `json:"FromNumber"`
	ToNumber    int64               `json:"ToNumber"`
	TanNumber   int64               `json:"TanNumber"`
	FromTrade   map[string]Trade    `json:"FromTrade"`
	ToTrade     map[string]Trade    `json:"ToTrade"`
	TanReport   map[int64]TanReport `json:"TanReport"`
}

/*
审核相关
*/
type Examine struct {
	IsExamine   string `json:"IsExamine"`
	ExamineType string `json:"ExamineType"`
	Examiner    string `json:"Examiner"`
	Sign        string `json:"Sign"`
}

/*
公司基本信息
*/
type CompanyInfo struct {
	Name                     string                   `json:"Name"`
	Type                     string                   `json:"Type"`
	Owner                    string                   `json:"Owner"`
	RegistrationNumber       string                   `json:"RegistrationNumber"` // 统一社会信用代码
	Address                  string                   `json:"Address"`
	BusinessScope            string                   `json:"BusinessScope"`       // 经营范围
	Contact                  Contact                  `json:"Contact"`             // 联系方式
	EstablishmentDate        string                   `json:"EstablishmentDate"`   // 成立日期
	RegisteredCapital        string                   `json:"RegisteredCapital"`   // 注册资本
	TaxRegistration          string                   `json:"TaxRegistration"`     // 税务登记证明文件路径或ID
	OrganizationCode         string                   `json:"OrganizationCode"`    // 组织机构代码证文件路径或ID
	BusinessLicense          string                   `json:"BusinessLicense"`     // 营业执照文件路径或ID
	CertificationStatus      string                   `json:"CertificationStatus"` // 认证状态，如是否通过环境认证
	AuthorizedRepresentative AuthorizedRepresentative `json:"AuthorizedRepresentative"`
	Status                   string                   `json:"Status"`
}

type AuthorizedRepresentative struct {
	Name             string `json:"Name"`
	Position         string `json:"Position"`
	IDNumber         string `json:"IDNumber"`
	AuthorizationDoc string `json:"AuthorizationDoc"` // 授权代表授权书文件路径或ID
}

type Contact struct {
	Phone string `json:"Phone"`
	Email string `json:"Email"`
}

/*
订单信息
*/
type Trade struct {
	TradeId     string  `json:"TradeId"`
	FromAccount string  `json:"FromAccount"`
	ToAccount   string  `json:"ToAccount"`
	Volume      float64 `json:"Volume"`
	Price       float64 `json:"Price"`
}

type TradeMap struct {
	Number int64            `json:"Number"`
	Trade  map[string]Trade `json:"Trade"`
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
	case "tanReportRegister":
		return t.TanReportRegister(stub, args)

	default:
		return shim.Error("Unsupported function")
	}
	return shim.Success(nil)
}

/*
创建一个account对应的User存储结构体并实现上链
*/
func (t *SimpleAsset) UserRegister(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of args.Expecting 5")
	}
	acc := args[0]
	info := args[1]
	bal := args[2]
	if acc == "" || bal == "" || info == "" {
		return shim.Error("Invalid args.")
	}
	accountByes, err := stub.GetState(acc)
	if err == nil && len(accountByes) != 0 {
		return shim.Error("account already exists")
	}
	balance, _ := strconv.ParseFloat(bal, 10)
	var companyInfo CompanyInfo
	err = json.Unmarshal([]byte(info), &companyInfo)
	if err != nil {
		return shim.Error("can't change")
	}
	user := User{
		Account:     acc,
		CompanyInfo: companyInfo,
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
	var Trademap TradeMap
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
	tradeByes, err := stub.GetState("TradeMap")
	if err != nil {
		return shim.Error("Search error!!")
	}
	if len(tradeByes) == 0 {
		Trademap.Trade = make(map[string]Trade)
	}
	err = json.Unmarshal(tradeByes, &Trademap)
	if err != nil {
		return shim.Error("can't change")
	}
	fromByes, err := stub.GetState(from)
	if err != nil && len(fromByes) == 0 {
		return shim.Error("Search error or no from user!!")
	}
	var fromUser User
	err = json.Unmarshal(fromByes, &fromUser)
	if err != nil {
		return shim.Error("can't change")
	}
	if fromUser.FromNumber == 0 {
		fromUser.FromTrade = make(map[string]Trade)
	}
	volume, _ := strconv.ParseFloat(vol, 10)
	price, _ := strconv.ParseFloat(pri, 10)
	trade := Trade{
		TradeId:     id,
		FromAccount: from,
		Volume:      volume,
		Price:       price,
	}
	fromUser.FromTrade[id] = trade
	fromUser.FromNumber = fromUser.FromNumber + 1
	Trademap.Trade[id] = trade
	Trademap.Number = Trademap.Number + 1
	traByes, err := json.Marshal(Trademap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("TradeMap", traByes); err != nil {
		return shim.Error("Failed to put state")
	}
	froByes, err := json.Marshal(fromUser)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(from, froByes); err != nil {
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
	var Trademap TradeMap
	id := args[0]
	TradeBytes, err := stub.GetState("TradeMap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if len(TradeBytes) == 0 {
		return shim.Error("Wrong !!")
	}
	err = json.Unmarshal(TradeBytes, &Trademap)
	if err != nil {
		return shim.Error("can't change")
	}
	idBytes, err := json.Marshal(Trademap.Trade[id])
	if err != nil {
		return shim.Error("error!")
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
	var Trademap TradeMap
	id := args[0]
	TradeBytes, err := stub.GetState("TradeMap")
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if len(TradeBytes) == 0 {
		return shim.Error("Wrong !!")
	}
	err = json.Unmarshal(TradeBytes, &Trademap)
	if err != nil {
		return shim.Error("can't change")
	}
	delete(Trademap.Trade, id)
	Trademap.Number = Trademap.Number - 1
	traByes, err := json.Marshal(Trademap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("TradeMap", traByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

/*
执行交易from->to
*/
func (t *SimpleAsset) Transaction(stub shim.ChaincodeStubInterface, args []string) peer.Response { //执行资金的密态转移
	if len(args) != 2 {
		return shim.Error("Incorrect number of args.Expecting 5")
	}
	var from, to User
	var Trademap TradeMap
	id, userTo := args[0], args[1]
	tradeByes, err := stub.GetState("TradeMap")
	if err != nil {
		return shim.Error("Failed to get tradeMap state")
	}
	if err = json.Unmarshal(tradeByes, &Trademap); err != nil {
		return shim.Error("Failed to unmarshal userFrom")
	}
	existFrom, err := stub.GetState(Trademap.Trade[id].FromAccount)
	if err == nil && len(existFrom) == 0 {
		return shim.Error("sender does not exist")
	}
	existTo, err := stub.GetState(userTo)
	if err == nil && len(existTo) == 0 {
		return shim.Error("receiver does not exist")
	}

	if id == "" || userTo == "" {
		return shim.Error("Invalid args")
	}
	price := Trademap.Trade[id].Price
	volume := Trademap.Trade[id].Volume
	userFromBytes, err := stub.GetState(Trademap.Trade[id].FromAccount)
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

	from.Balance = from.Balance - price*volume
	to.Balance = to.Balance + price*volume
	if from.FromNumber == 0 {
		from.FromTrade = make(map[string]Trade)
	}
	if to.ToNumber == 0 {
		to.ToTrade = make(map[string]Trade)
	}
	Trademap.Trade[id] = Trade{
		TradeId:     Trademap.Trade[id].TradeId,
		FromAccount: Trademap.Trade[id].FromAccount,
		ToAccount:   userTo,
		Price:       Trademap.Trade[id].Price,
		Volume:      Trademap.Trade[id].Volume,
	}
	from.FromTrade[id] = Trademap.Trade[id]
	from.FromNumber = from.FromNumber + 1
	to.ToTrade[id] = Trademap.Trade[id]
	to.ToNumber = to.ToNumber + 1
	delete(Trademap.Trade, id)
	Trademap.Number = Trademap.Number - 1
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
	traByes, err := json.Marshal(Trademap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("TradeMap", traByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

/*
添加碳报告
*/
func (t *SimpleAsset) TanReportRegister(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	acc := args[0]
	tanReport := args[1]
	idBytes, err := stub.GetState(acc)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	var user User
	err = json.Unmarshal(idBytes, &user)
	if user.TanNumber == 0 {
		user.TanReport = make(map[int64]TanReport)
	}
	if err != nil {
		return shim.Error("Error 2 !!")
	}
	var Tanreport TanReport
	err = json.Unmarshal([]byte(tanReport), &Tanreport)
	if err != nil {
		return shim.Error("Error 3 !!")
	}
	user.TanReport[user.TanNumber] = Tanreport
	user.TanNumber = user.TanNumber + 1
	newUser, err := json.Marshal(user)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(user.Account, newUser); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
