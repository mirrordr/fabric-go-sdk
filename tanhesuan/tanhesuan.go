package tanhesuan

// 化石燃料燃烧 中英文转换资料由化工百科https://www.chembk.com/搜索得到
type Fossil_Fuel_Combustion struct {
	Anthracite                                   float64 `json:"Anthracite"`                                   //无烟煤
	Bituminous_coal                              float64 `json:"Bituminous_Coal"`                              //烟煤
	Brown_coal                                   float64 `json:"Brown_Coal"`                                   //褐煤
	Briquette                                    float64 `json:"Briquette"`                                    //型煤
	Coke                                         float64 `json:"Coke"`                                         //焦炭
	Crude                                        float64 `json:"Crude"`                                        //原油
	Fuel_oil                                     float64 `json:"Fuel_Oil"`                                     //燃料油
	Gasoline                                     float64 `json:"Gasoline"`                                     //汽油
	Diesel_fuel                                  float64 `json:"Diesel_Fuel"`                                  //柴油
	General_kerosene                             float64 `json:"General_Kerosene"`                             //一般煤油
	Aviation_gasoline                            float64 `json:"Aviation_Gasoline"`                            //航空汽油
	Aviation_kerosene                            float64 `json:"Aviation_kerosene"`                            //航空煤油
	Liquefied_natural_gas                        float64 `json:"Liquefied_natural_gas"`                        //液化天然气
	Liquefied_petroleum_gas                      float64 `json:"Liquefied_petroleum_gas"`                      //液化石油气
	Petroleum_products                           float64 `json:"Petroleum_products"`                           //炼厂干气
	Naphtha                                      float64 `json:"Naphtha"`                                      //石脑油
	Petroleum_coke                               float64 `json:"Petroleum_Coke"`                               //石油焦
	Methane                                      float64 `json:"Methane"`                                      //天然气
	Coke_oven_gas                                float64 `json:"Coke_oven_gas"`                                //焦炉煤气
	Coal_tar                                     float64 `json:"Coal_tar"`                                     //煤焦油、焦油
	Water_gas                                    float64 `json:"Water_gas"`                                    //水煤气
	Clenedcoal                                   float64 `json:"Clenedcoal"`                                   //洗精煤
	Washing_middings                             float64 `json:"Washing_middings"`                             //洗中煤
	Slime                                        float64 `json:"Slime"`                                        //煤泥
	Producer_gas                                 float64 `json:"Producer_gas"`                                 //发生炉煤气
	Catalytic_cracking_process_gas_of_heavy_oil  float64 `json:"Catalytic_cracking_process_gas_of_heavy_oil"`  //重油催化裂解煤气
	Catalytic_pyrolysis_process_gas_of_heavy_oil float64 `json:"Catalytic_pyrolysis_process_gas_of_heavy_oil"` //重油催化热煤气
	Coke_gas                                     float64 `json:"Coke_gas"`                                     //焦炭制气
	High_pressure_gasification_gas               float64 `json:"High_pressure_gasification_gas"`               //压力气化煤气
	Pure_coke                                    float64 `json:"Pure_coke"`                                    //蓝炭
	Blastfurnace_gas                             float64 `json:"Blastfurnace_gas"`                             //高炉煤气
	Converter_gas                                float64 `json:"Converter_gas"`                                //转炉煤气
	Semicoke_gas                                 float64 `json:"Semicoke_gas"`                                 //半焦气
}

// 电力热力排放
type Electricity_And_Heat_Emissions struct {
	Electricity float64 `json:"Electricity"` //电力净购入量
	Heat        float64 `json:"Heat"`        //热力净购入量
	Region      string  `json:"Region"`      //地区，用于核算电力因子
}

// 陶瓷行业工业生产过程碳排放
type Ceramics_Indusry_Production_Process struct {
	Material_weight            float64 `json:"Material_weight"`            //原料使用量
	Material_utilization_ratio float64 `json:"Material_utilization_ratio"` //原料利用率
	CaCO3_content              float64 `json:"CaCO3_content"`              //碳酸钙含量
	MgCO3_content              float64 `json:"MgCO3_content"`              //碳酸镁含量
}

// 水泥行业替代燃料或废弃物种非生物质碳的燃烧排放
type Cement_Industry_Alternative_Fuels_Burning struct {
	Waste_oil      float64 `json:"Waste_oil"`      //废油
	Waste_tires    float64 `json:"Waste_tires"`    //废轮胎
	Plastics       float64 `json:"Plastics"`       //塑料
	Spent_solvents float64 `json:"Spent_solvents"` //废溶剂
	Waste_leather  float64 `json:"Waste_leather"`  //废皮革
	Waste_FRP      float64 `json:"Waste_FRP"`      //废玻璃钢
}

// 水泥行业工业生产过程的排放
type Cement_Industry_Production_Process struct {
	Clinker_yield           float64 `json:"Clinker_yield"`           //熟料产量
	CaO_content             float64 `json:"CaO_content"`             //熟料中碳酸盐分解的氧化钙含量
	MgO_content             float64 `json:"MgO_content"`             //熟料中碳酸盐分解的氧化镁含量
	Kiln_head_dust_weight   float64 `json:"Kiln_head_dust_weight"`   //窑头粉尘重量
	Bypass_vent_dust_weight float64 `json:"Bypass_vent_dust_weight"` //旁路放风粉尘重量
	Raw_meal_weight         float64 `json:"Raw_meal_weight"`         //生料重量
	Nonfuel_carbon_content  float64 `json:"Nonfuel_carbon_content"`  //生料中非燃料碳含量
}

// 平板玻璃产业替代燃料或废弃物种非生物质碳的燃烧排放
type Flat_glass_Industry_Alternative_Fuels_Burning struct {
	Toner_consumption float64 `json:"Toner_consumption"` //碳粉消耗量
	Carbon_content    float64 `json:"Carbon_content"`    //碳粉含碳量
}

// 平板玻璃产业原料分解过程产生的排放
type Cement_Industry_Decomposition_Process struct {
	Limestone_weight     float64 `json:"Limestone_weight"`     //石灰石重量
	Limestone_ratio      float64 `json:"Limestone_ratio"`      //石灰石煅烧比例
	Dolomite_weight      float64 `json:"Dolomite_weight"`      //白云石重量
	Dolomite_ratio       float64 `json:"Dolomite_ratio"`       //白云石煅烧比例
	Soda_ash_weight      float64 `json:"Soda_ash_weight"`      //纯碱重量
	Soda_ash_ratio       float64 `json:"Soda_ash_ratio"`       //纯碱煅烧比例
	Magnesite_weight     float64 `json:"Magnesite_weight"`     //菱镁石重量
	Magnesite_ratio      float64 `json:"Magnesite_ratio"`      //菱镁石煅烧比例
	Ankerite_weight      float64 `json:"Ankerite_weight"`      //铁白云石重量
	Ankerite_ratio       float64 `json:"Ankerite_ratio"`       //铁白云石煅烧比例
	Siderite_weight      float64 `json:"Siderite_weight"`      //菱铁矿重量
	Siderite_ratio       float64 `json:"Siderite_ratio"`       //菱铁矿煅烧比例
	Rhodochrosite_weight float64 `json:"Rhodochrosite_weight"` //菱锰矿重量
	Rhodochrosite_ratio  float64 `json:"Rhodochrosite_ratio"`  //菱锰矿煅烧比例
}

// 镁冶炼行业工业生产过程、能源作为原材料用途的排放
type Magnesium_smelting_Industry_Production_Process struct {
	Ferrosilicon_yield   float64 `json:"Ferrosilicon_yield"`   //硅铁产量
	Dolomite_consumption float64 `json:"Dolomite_consumption"` //白云石作为原料的消耗量
}

// 发电企业脱硫
type power_generation_Industry_DeS_Process struct {
	Caco3  float64 `json:"CaCO3"`
	Mgco3  float64 `json:"MgCO3"`
	Na2co3 float64 `json:"Na2CO3"`
	Baco3  float64 `json:"BaCO3"`
	Li2co3 float64 `json:"Li2CO3"`
	K2co3  float64 `json:"K2CO3"`
	Srco3  float64 `json:"SrCO3"`
	Nahco3 float64 `json:"NahCO3"`
	Feco3  float64 `json:"FeCO3"`
}

// 电网企业Esf6
type grid_Industry_Esf6 struct {
	//
}

// 电网企业网损排放
type grid_Industry_trans_losses struct {
	EL_sold    float64 `json:"EL_sold"`    //售电量，即终端用户用电量
	EL_feed_in float64 `json:"EL_feed_in"` //电厂上网电量
	EL_input   float64 `json:"EL_input"`   //自外省输入电量
	EL_output  float64 `json:"EL_output"`  //向外省输出电量
}

// 钢铁企业工业生产排放
type steel_Industry_production_E struct {
	Imestone              float64 `json:"Imestone"`              //石灰石
	Dolomite              float64 `json:"Dolomite"`              //白云石
	Electrodes            float64 `json:"Electrodes"`            //电极
	Raw_iron              float64 `json:"Raw_iron"`              //生铁
	Directly_reduced_iron float64 `json:"Directly_reduced_iron"` //直接还原铁
	Nickel_iron           float64 `json:"Nickel_iron"`           //镍铁合金
	Ferrochrome           float64 `json:"Ferrochrome"`           //铬铁合金
	Molybdenum_iron       float64 `json:"Molybdenum_iron"`       //钼铁合金
}

// 钢铁企业固碳
type steel_Industry_sequestration_E struct {
	Raw_iron    float64 `json:"Raw_iron"`    //生铁
	Crude_steel float64 `json:"Crude_steel"` //粗钢
	Wood_spirit float64 `json:"Wood_spirit"` //甲醇
	Other_seq   float64 `json:"Other_seq"`   //其他固碳产品或副产品
}

// 化学企业生产过程
type chem_Industry_E struct {
	Anthracite           float64 `json:"Anthracite"`           //无烟煤
	Coke                 float64 `json:"Coke"`                 //焦炭
	Crude                float64 `json:"Crude"`                //原油
	Naphtha              float64 `json:"Naphtha"`              //石脑油
	Petroleum_coke       float64 `json:"Petroleum_coke"`       //石油焦
	Carbon_electrode     float64 `json:"Carbon_electrode"`     //碳电极
	Methane              float64 `json:"Methane"`              //天然气
	Ethylene             float64 `json:"Ethylene"`             //乙烯
	Propylene            float64 `json:"Propylene"`            //丙烯
	Niaosu               float64 `json:"Niaosu"`               //尿素
	Ammonium_bicarbonate float64 `json:"Ammonium_bicarbonate"` //碳酸氢铵
	Wood_spirit          float64 `json:"Wood_spirit"`          //甲醇
	Calcium_carbide      float64 `json:"Calcium_carbide"`      //电石
	Stove_ash            float64 `json:"Stove_ash"`            //炉渣
	Dust                 float64 `json:"Dust"`                 //粉尘
	Sluge                float64 `json:"Sluge"`                //污泥
	Limestone            float64 `json:"Limestone"`            //石灰石
	Dolomite             float64 `json:"Dolomite"`             //白云石
	Magnesium_oxide      float64 `json:"Magnesium_oxide"`      //菱镁石
	Clays                float64 `json:"Clays"`                // 粘土
	// 以下为各硝酸生产工艺对应的硝酸产量
	Hyperbaric  float64 `json:"Hyperbaric"`  //高压
	Medium_pres float64 `json:"Medium_pres"` //中压
	Atmospheric float64 `json:"Atmospheric"` //常压
	Double_pres float64 `json:"Double_pres"` //双加压
	Synthetic   float64 `json:"Synthetic"`   //综合法
	Low_pres    float64 `json:"Low_pres"`    //低压法
	//乙二酸生产类型对应的乙二酸产量
	N_acid_oxi    float64 `json:"N_acid_oxi"`    //硝酸氧化
	Other_methods float64 `json:"Other_methods"` //其他
	Co2_rec       float64 `json:"Co2_Rec"`       //Co2回收量
}

// 电解铝企业
type aluminum_electrolysis_industry_E struct {
	Primary_aluminum1          float64 `json:"Primary_aluminum1"`          //原铝产量
	Primary_aluminum2          float64 `json:"Primary_aluminum2"`          //原铝产量
	Limestone_cons             float64 `json:"Limestone_cons"`             //石灰石原料消耗量
	Net_consumption            float64 `json:"Net_consumption"`            //吨铝炭阳极净耗
	Average_sulfur_content     float64 `json:"Average_sulfur_content"`     //炭阳极平均含硫量
	Average_ash_content        float64 `json:"Average_ash_content"`        //	炭阳极平均灰分含量
	CF4                        float64 `json:"CF4"`                        //阳极效应的 CF4 排放因子
	C2F6                       float64 `json:"C2F6"`                       //阳极效应的 C2F6 排放因子
	Average_daily_anode_effect float64 `json:"Average_daily_anode_effect"` //平均每天每槽阳极效应持续时间
	Limestone_factors          float64 `json:"Limestone_factors"`          //煅烧石灰石的排放因子
}

var co2paifangyingzi = map[string]float64{
	"LN":   0.578,
	"JL":   0.564,
	"HLJ":  0.654,
	"BJ":   0.595,
	"TJ":   0.688,
	"HB":   0.736,
	"SX":   0.707,
	"NMG":  0.8,
	"SD":   0.546,
	"SH":   0.333,
	"JS":   0.601,
	"ZJ":   0.418,
	"AH":   0.755,
	"FJ":   0.363,
	"JX":   0.474,
	"HN":   0.599,
	"HUB":  0.31,
	"HUN":  0.453,
	"CQ":   0.363,
	"SC":   0.104,
	"GD":   0.369,
	"GX":   0.336,
	"HAIN": 0.326,
	"GZ":   0.398,
	"YN":   0.1,
	"SHX":  0.607,
	"GS":   0.443,
	"QH":   0.067,
	"NX":   0.724,
	"XJ":   0.720,
}
