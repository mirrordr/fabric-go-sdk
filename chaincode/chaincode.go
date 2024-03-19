package main

import (
	"chaincode/tanhesuan"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type SimpleAsset struct {
}

type TanReport struct {
	Huashi        tanhesuan.Fossil_Fuel_Combustion                         `json:"Huashi"`        //化石燃料信息
	Taocizhuanyou tanhesuan.Ceramics_Indusry_Production_Process            `json:"Taocizhuanyou"` //陶瓷企业的专有信息
	Dianli        tanhesuan.Electricity_And_Heat_Emissions                 `json:"Dianli"`        //电力热力信息
	Ma            tanhesuan.Magnesium_smelting_Industry_Production_Process `json:"Ma"`            //镁工业特有
	Time          time.Time                                                `json:"Time"`          //提交时间
	Final         float64                                                  `json:"Final"`         //最终结果
	Type          string                                                   `json:"Type"`          //碳报告类型
	Examine       Examine                                                  `json:"Examine"`       //监管签名，只有签名了的才可以用于碳币的生成和交易等
}

type TaociHeyunsuan struct {
	Huashimodel1 tanhesuan.Fossil_Fuel_Combustion `json:"Huashimodel1"`
	Huashimodel2 tanhesuan.Fossil_Fuel_Combustion `json:"Huashimodel2"`
	Huashimodel3 tanhesuan.Fossil_Fuel_Combustion `json:"Huashimodel3"`
}

type MgHeyunsuan struct {
	Huashimodel1 tanhesuan.Fossil_Fuel_Combustion                         `json:"Huashimodel1"`
	Huashimodel2 tanhesuan.Fossil_Fuel_Combustion                         `json:"Huashimodel2"`
	Huashimodel3 tanhesuan.Fossil_Fuel_Combustion                         `json:"Huashimodel3"`
	Mg           tanhesuan.Magnesium_smelting_Industry_Production_Process `json:"Mg"`
}

/*
保存用户
*/
type User struct {
	Account       string              `json:"Account"`       //用户的账号+账号密码的hash
	CompanyInfo   CompanyInfo         `json:"CompanyInfo"`   //公司信息
	Balance       float64             `json:"Balance"`       //用户余额
	Volume        float64             `json:"Volume"`        //公司碳额度
	Examine       Examine             `json:"Examine"`       //审核签名（如果这栏为空代表没有审核显示没有审核通过）
	FromNumber    int64               `json:"FromNumber"`    //发起交易的数量
	ToNumber      int64               `json:"ToNumber"`      //选择交易的数量
	TanNumber     int64               `json:"TanNumber"`     //上传碳报告的数量
	ProceedNumber int64               `json:"ProceedNumber"` //诉讼数量
	FromTrade     map[string]Trade    `json:"FromTrade"`     //发起交易的交易信息
	ToTrade       map[string]Trade    `json:"ToTrade"`       //选择交易的交易信息
	TanReport     map[int64]TanReport `json:"TanReport"`     //谈报告的具体信息
	Proceed       map[string]Proceed  `json:"Proceed"`
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
	TradeId     string  `json:"TradeId"`     //ID
	FromAccount string  `json:"FromAccount"` //From
	ToAccount   string  `json:"ToAccount"`   //To
	Volume      float64 `json:"Volume"`      //交易量
	Price       float64 `json:"Price"`       //交易单价
}

type TradeMap struct {
	Number int64            `json:"Number"` //总共的提出交易的数量
	Trade  map[string]Trade `json:"Trade"`  //请求交易的Map
}

type TanReportMap struct {
	Number    int64                `json:"Number"`
	TanReport map[string]TanReport `json:"TanReport"`
}

type Data_auditor struct {
	ExaminePK *big.Int `json:"ExaminePK"`
	ExamineSK *big.Int `json:"ExamineSK"`
}

type Proceed struct {
	PrID   string  `json:"PrID"`
	ID     string  `json:"TradeID"`
	From   string  `json:"From"`
	To     string  `json:"To"`
	Price  float64 `json:"Price"`
	Volume float64 `json:"Volume"`
	Final  string  `json:"Final"`
}

type ProceedMap struct {
	Number  int64              `json:"Number"`
	Proceed map[string]Proceed `json:"Proceed"`
}

type ED struct {
	Taoci float64 `json:"陶瓷"`
	Mg    float64 `json:"镁"`
}

// Init /*区块链的初始化
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	test1 := User{
		Account: "test1",
		CompanyInfo: CompanyInfo{
			Name: "test1Company",
			Type: "陶瓷",
		},
		Balance: 100,
		Volume:  100,
	}
	test1Byes, err := json.Marshal(test1)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(test1.Account, test1Byes); err != nil {
		return shim.Error("Failed to put state")
	}
	test2 := User{
		Account: "test2",
		CompanyInfo: CompanyInfo{
			Name: "test2Company",
			Type: "镁",
		},
		Balance: 100,
		Volume:  100,
	}
	test2Byes, err := json.Marshal(test2)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(test2.Account, test2Byes); err != nil {
		return shim.Error("Failed to put state")
	}
	tradeMap := TradeMap{
		Number: 0,
		Trade:  make(map[string]Trade),
	}
	tradeByes, err := json.Marshal(tradeMap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("TradeMap", tradeByes); err != nil {
		return shim.Error("Failed to put state")
	}
	Taocimodel1 := tanhesuan.Fossil_Fuel_Combustion{
		Anthracite:              23.2,
		Bituminous_coal:         22.3,
		Brown_coal:              14.8,
		Briquette:               17.5,
		Coke:                    28.4,
		Crude:                   41.8,
		Gasoline:                43.1,
		Diesel_fuel:             42.7,
		General_kerosene:        43.1,
		Fuel_oil:                41.8,
		Coal_tar:                33.5,
		Liquefied_natural_gas:   51.4,
		Liquefied_petroleum_gas: 50.2,
		Methane:                 389.3,
		Water_gas:               10.4,
		Coke_oven_gas:           173.5,
	}
	Taocimodel2 := tanhesuan.Fossil_Fuel_Combustion{
		Anthracite:              27.8,
		Bituminous_coal:         25.6,
		Brown_coal:              27.8,
		Briquette:               33.6,
		Coke:                    28.8,
		Crude:                   20.1,
		Gasoline:                18.9,
		Diesel_fuel:             20.2,
		General_kerosene:        19.6,
		Fuel_oil:                21,
		Coal_tar:                22,
		Liquefied_natural_gas:   15.3,
		Liquefied_petroleum_gas: 17.2,
		Methane:                 15.3,
		Water_gas:               12.2,
		Coke_oven_gas:           13.6,
	}
	Taocimodel3 := tanhesuan.Fossil_Fuel_Combustion{
		Anthracite:              0.94,
		Bituminous_coal:         0.93,
		Brown_coal:              0.96,
		Briquette:               0.9,
		Coke:                    0.93,
		Crude:                   0.98,
		Gasoline:                0.98,
		Diesel_fuel:             0.98,
		General_kerosene:        0.98,
		Fuel_oil:                0.98,
		Coal_tar:                0.98,
		Liquefied_natural_gas:   0.98,
		Liquefied_petroleum_gas: 0.98,
		Methane:                 0.99,
		Water_gas:               0.99,
		Coke_oven_gas:           0.99,
	}
	Taoci := TaociHeyunsuan{
		Huashimodel1: Taocimodel1,
		Huashimodel2: Taocimodel2,
		Huashimodel3: Taocimodel3,
	}
	taociByes, err := json.Marshal(Taoci)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("Taoci", taociByes); err != nil {
		return shim.Error("Failed to put state")
	}
	mgmodel1 := tanhesuan.Fossil_Fuel_Combustion{
		Anthracite:              20.304,
		Bituminous_coal:         19.57,
		Brown_coal:              14.08,
		Clenedcoal:              26.344,
		Pure_coke:               28.435,
		Coke:                    28.447,
		Crude:                   41.816,
		Fuel_oil:                41.816,
		Gasoline:                43.07,
		Diesel_fuel:             42.652,
		General_kerosene:        44.75,
		Liquefied_natural_gas:   41.868,
		Liquefied_petroleum_gas: 50.176,
		Coal_tar:                33.453,
		Coke_oven_gas:           173.54,
		Blastfurnace_gas:        33,
		Converter_gas:           84,
		Producer_gas:            52.27,
		Methane:                 389.31,
		Semicoke_gas:            81,
		Petroleum_products:      45.998,
	}
	mgmodel2 := tanhesuan.Fossil_Fuel_Combustion{
		Anthracite:              27.49,
		Bituminous_coal:         26.18,
		Brown_coal:              28,
		Clenedcoal:              25.4,
		Pure_coke:               29.42,
		Coke:                    29.5,
		Crude:                   20.1,
		Fuel_oil:                21.1,
		Gasoline:                18.9,
		Diesel_fuel:             20.2,
		General_kerosene:        19.6,
		Liquefied_natural_gas:   17.2,
		Liquefied_petroleum_gas: 17.2,
		Coal_tar:                22,
		Coke_oven_gas:           12.1,
		Blastfurnace_gas:        70.8,
		Converter_gas:           49.6,
		Producer_gas:            12.2,
		Methane:                 15.3,
		Semicoke_gas:            11.96,
		Petroleum_products:      18.2,
	}
	mgmodel3 := tanhesuan.Fossil_Fuel_Combustion{
		Anthracite:              0.94,
		Bituminous_coal:         0.93,
		Brown_coal:              0.96,
		Clenedcoal:              0.9,
		Pure_coke:               0.93,
		Coke:                    0.93,
		Crude:                   0.98,
		Fuel_oil:                0.98,
		Gasoline:                0.98,
		Diesel_fuel:             0.98,
		General_kerosene:        0.98,
		Liquefied_natural_gas:   0.98,
		Liquefied_petroleum_gas: 0.98,
		Coal_tar:                0.98,
		Coke_oven_gas:           0.99,
		Blastfurnace_gas:        0.99,
		Converter_gas:           0.99,
		Producer_gas:            0.99,
		Methane:                 0.99,
		Semicoke_gas:            0.99,
		Petroleum_products:      0.99,
	}
	mgmodel4 := tanhesuan.Magnesium_smelting_Industry_Production_Process{
		Ferrosilicon_yield:   2.79,
		Dolomite_consumption: 0.98,
	}
	mg := MgHeyunsuan{
		Huashimodel1: mgmodel1,
		Huashimodel2: mgmodel2,
		Huashimodel3: mgmodel3,
		Mg:           mgmodel4,
	}
	mgByes, err := json.Marshal(mg)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("Mg", mgByes); err != nil {
		return shim.Error("Failed to put state")
	}
	tanReportMap := TanReportMap{
		Number:    0,
		TanReport: make(map[string]TanReport),
	}
	tanmapByes, err := json.Marshal(tanReportMap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("TanReportMap", tanmapByes); err != nil {
		return shim.Error("Failed to put state")
	}
	proceedMap := ProceedMap{
		Number:  0,
		Proceed: make(map[string]Proceed),
	}
	promapByes, err := json.Marshal(proceedMap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("ProceedMap", promapByes); err != nil {
		return shim.Error("Failed to put state")
	}
	ed := ED{
		Taoci: 5000,
		Mg:    5000,
	}
	edByes, err := json.Marshal(ed)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("ED", edByes); err != nil {
		return shim.Error("Failed to put state")
	}
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
	case "changeTaoci":
		return t.ChangeTaoci(stub, args)
	case "changeMg":
		return t.ChangeMg(stub, args)
	case "tanHesuan":
		return t.TanHesuan(stub, args)
	case "proceedRegister":
		return t.ProceedRegister(stub, args)
	case "proceed":
		return t.Proceed(stub, args)
	case "changeED":
		return t.ChangeED(stub, args)

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
	info1 := args[1]
	bal := args[2]
	info := strings.Replace(info1, "\\", "", -1)
	fmt.Println(info1)
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
		Volume:      100,
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
	err = json.Unmarshal(tradeByes, &Trademap)
	if err != nil {
		return shim.Error("can't change 1")
	}
	fromByes, err := stub.GetState(from)
	if err != nil && len(fromByes) == 0 {
		return shim.Error("Search error or no from user!!")
	}
	var fromUser User
	err = json.Unmarshal(fromByes, &fromUser)
	if err != nil {
		return shim.Error("can't change 2")
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
	if from.Balance < price*volume {
		return shim.Error("no balance")
	}
	if to.Volume < volume {
		return shim.Error("no volume")
	}
	from.Balance = from.Balance + price*volume
	to.Balance = to.Balance - price*volume
	from.Volume = from.Volume - volume
	to.Volume = to.Volume + volume
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
	var TanreportMap TanReportMap
	mapBytes, err := stub.GetState("TanReportMap")
	err = json.Unmarshal(mapBytes, &TanreportMap)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	var Tanreport TanReport
	err = json.Unmarshal([]byte(tanReport), &Tanreport)
	if err != nil {
		return shim.Error("Error 3 !!")
	}
	Time := time.Now()
	Tanreport.Time = Time
	Tanreport.Type = user.CompanyInfo.Type
	user.TanReport[user.TanNumber] = Tanreport
	TanreportMap.Number = TanreportMap.Number + 1
	if _, ok := TanreportMap.TanReport[user.Account]; ok {
		return shim.Error("已经有未审核通过的碳核算报告")
	}
	TanreportMap.TanReport[user.Account] = Tanreport
	user.TanNumber = user.TanNumber + 1
	newUser, err := json.Marshal(user)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(user.Account, newUser); err != nil {
		return shim.Error("Failed to put state")
	}
	newTan, err := json.Marshal(TanreportMap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("TanReportMap", newTan); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

func (t *SimpleAsset) ChangeTaoci(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	taoci := args[0]
	var newTaoci TaociHeyunsuan
	err := json.Unmarshal([]byte(taoci), &newTaoci)
	if err != nil {
		return shim.Error("can't change 3")
	}
	taociByes, err := json.Marshal(newTaoci)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("Taoci", taociByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

func (t *SimpleAsset) ChangeMg(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	mg := args[0]
	var newmg MgHeyunsuan
	err := json.Unmarshal([]byte(mg), &newmg)
	if err != nil {
		return shim.Error("can't change 3")
	}
	mgByes, err := json.Marshal(newmg)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("Mg", mgByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

func (t *SimpleAsset) TanHesuan(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	acc := args[0]
	finally := args[1]
	var TanreportMap TanReportMap
	mapBytes, err := stub.GetState("TanReportMap")
	err = json.Unmarshal(mapBytes, &TanreportMap)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	var Taoci TaociHeyunsuan
	taociBytes, err := stub.GetState("Taoci")
	err = json.Unmarshal(taociBytes, &Taoci)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	var ed ED
	edBytes, err := stub.GetState("ED")
	err = json.Unmarshal(edBytes, &ed)
	if err != nil {
		return shim.Error("Failed to get state")
	}

	var Mg MgHeyunsuan
	mgBytes, err := stub.GetState("Mg")
	err = json.Unmarshal(mgBytes, &Mg)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	var report TanReport
	var edd float64
	if TanreportMap.TanReport[acc].Type == "陶瓷" {
		edd = ed.Taoci
	} else {
		edd = ed.Mg
	}
	report = TanreportMap.TanReport[acc]
	switch report.Type {
	case "陶瓷":
		report.Final = tanhesuan.Taoci(&report.Huashi, &report.Taocizhuanyou, &report.Dianli, Taoci.Huashimodel1, Taoci.Huashimodel2, Taoci.Huashimodel3)
		break
	case "镁":
		report.Final = tanhesuan.Mayanlian(&report.Huashi, &report.Ma, &report.Dianli, Mg.Huashimodel1, Mg.Huashimodel2, Mg.Huashimodel3, Mg.Mg)
		break
	default:
		return shim.Error("No type")
	}
	TanreportMap.TanReport[acc] = report
	idBytes, err := stub.GetState(acc)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	var user User
	err = json.Unmarshal(idBytes, &user)
	if err != nil {
		return shim.Error("Error 2 !!")
	}
	if finally == "False" {
		report.Final = -1
		edd = 1
	}
	user.TanReport[user.TanNumber-1] = report
	user.Volume = user.Volume + edd - report.Final

	delete(TanreportMap.TanReport, user.Account)
	TanreportMap.Number = TanreportMap.Number - 1

	newUser, err := json.Marshal(user)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState(user.Account, newUser); err != nil {
		return shim.Error("Failed to put state")
	}
	newTan, err := json.Marshal(TanreportMap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("TanReportMap", newTan); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

func (t *SimpleAsset) ProceedRegister(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of args.Expecting 3")
	}
	prid := args[0]
	id := args[1]
	userfrom := args[2]
	userto := args[3]
	var from, to User
	userFromBytes, err := stub.GetState(userfrom)
	if err != nil {
		return shim.Error("Failed to get userFrom state")
	}
	if err = json.Unmarshal(userFromBytes, &from); err != nil {
		return shim.Error("Failed to unmarshal userFrom")
	}
	userToByes, err := stub.GetState(userto)
	if err != nil {
		return shim.Error("Failed to get userFrom state")
	}
	if err = json.Unmarshal(userToByes, &to); err != nil {
		return shim.Error("Failed to unmarshal userFrom")
	}
	var proceed Proceed
	proceed = Proceed{
		PrID:   prid,
		ID:     id,
		From:   userfrom,
		To:     userto,
		Price:  from.FromTrade[id].Price,
		Volume: from.FromTrade[id].Volume,
	}
	var proceedMap ProceedMap
	proBytes, err := stub.GetState("ProceedMap")
	err = json.Unmarshal(proBytes, &proceedMap)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	for _, v := range proceedMap.Proceed {
		if v.ID == id {
			return shim.Error("have proceed")
		}
	}
	proceedMap.Proceed[prid] = proceed
	proceedMap.Number = proceedMap.Number + 1
	if from.ProceedNumber == 0 {
		from.Proceed = make(map[string]Proceed)
	}
	if to.ProceedNumber == 0 {
		to.Proceed = make(map[string]Proceed)
	}
	from.Proceed[prid] = proceed
	from.ProceedNumber = from.ProceedNumber + 1
	to.Proceed[prid] = proceed
	to.ProceedNumber = to.ProceedNumber + 1
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
	promapByes, err := json.Marshal(proceedMap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("ProceedMap", promapByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

func (t *SimpleAsset) Proceed(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of args.Expecting 2")
	}
	prid := args[0]
	fin := args[1]
	var proceedMap ProceedMap
	proBytes, err := stub.GetState("ProceedMap")
	err = json.Unmarshal(proBytes, &proceedMap)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	var proceed, newproceed Proceed
	proceed = proceedMap.Proceed[prid]
	var from, to User
	userFromBytes, err := stub.GetState(proceed.From)
	if err != nil {
		return shim.Error("Failed to get userFrom state")
	}
	if err = json.Unmarshal(userFromBytes, &from); err != nil {
		return shim.Error("Failed to unmarshal userFrom")
	}
	userToByes, err := stub.GetState(proceed.To)
	if err != nil {
		return shim.Error("Failed to get userFrom state")
	}
	if err = json.Unmarshal(userToByes, &to); err != nil {
		return shim.Error("Failed to unmarshal userFrom")
	}
	newproceed = Proceed{
		PrID:   prid,
		ID:     proceed.ID,
		From:   proceed.From,
		To:     proceed.To,
		Price:  proceed.Price,
		Volume: proceed.Volume,
		Final:  fin,
	}
	delete(proceedMap.Proceed, prid)
	proceedMap.Number = proceedMap.Number - 1
	from.Proceed[prid] = newproceed
	to.Proceed[prid] = newproceed
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
	promapByes, err := json.Marshal(proceedMap)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("ProceedMap", promapByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)

}

func (t *SimpleAsset) ChangeED(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of args.Expecting 1")
	}
	ed := args[0]
	var newED ED
	err := json.Unmarshal([]byte(ed), &newED)
	if err != nil {
		return shim.Error("can't change 3")
	}
	edByes, err := json.Marshal(newED)
	if err != nil {
		return shim.Error("marshal user error")
	}
	if err = stub.PutState("ED", edByes); err != nil {
		return shim.Error("Failed to put state")
	}
	return shim.Success(nil)
}

func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
