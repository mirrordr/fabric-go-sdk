package tanhesuan

import (
	"math/big"
	"reflect"
)

func Taoci(hushi *Fossil_Fuel_Combustion, taocizhuanyou *Ceramics_Indusry_Production_Process, dianli *Electricity_And_Heat_Emissions, model1 Fossil_Fuel_Combustion, model2 Fossil_Fuel_Combustion, model3 Fossil_Fuel_Combustion) (float64, float64, float64, float64) {
	var result float64
	StructFieldMot(&model1, &model2, &model3, hushi)
	a1 := SumFossil(&model1) * 44 / 12
	a2 := SumEle(dianli)
	a3 := taocizhuanyou.Material_weight * taocizhuanyou.Material_utilization_ratio * (44/100*taocizhuanyou.CaCO3_content + 44/84*taocizhuanyou.MgCO3_content)
	result = a1 + a2 + a3
	return result, a1, a2, a3
}

func Mayanlian(huashi *Fossil_Fuel_Combustion, ma *Magnesium_smelting_Industry_Production_Process, dianli *Electricity_And_Heat_Emissions, model1 Fossil_Fuel_Combustion, model2 Fossil_Fuel_Combustion, model3 Fossil_Fuel_Combustion, model4 Magnesium_smelting_Industry_Production_Process) (float64, float64, float64, float64) {
	var result float64
	StructFieldMot(&model1, &model2, &model3, huashi)
	a1 := SumFossil(&model1) * 44 / 12
	a2 := SumEle(dianli)
	StructFieldMot(&model4, ma)
	a3 := model4.Ferrosilicon_yield + model4.Dolomite_consumption*0.478
	result = a1 + a2 + a3
	return result, a1, a2, a3
}

func SumFossil(v1 *Fossil_Fuel_Combustion) float64 {
	var result float64
	result = v1.Anthracite + v1.Bituminous_coal + v1.Brown_coal + v1.Briquette + v1.Coke + v1.Crude + v1.Fuel_oil + v1.Gasoline + v1.Diesel_fuel + v1.General_kerosene + v1.Aviation_kerosene + v1.Aviation_gasoline + v1.Liquefied_natural_gas + v1.Liquefied_petroleum_gas + v1.Petroleum_products + v1.Naphtha + v1.Petroleum_coke + v1.Methane + v1.Coke_oven_gas + v1.Coal_tar + v1.Water_gas + v1.Clenedcoal + v1.Washing_middings + v1.Slime + v1.Producer_gas + v1.Catalytic_cracking_process_gas_of_heavy_oil + v1.Catalytic_pyrolysis_process_gas_of_heavy_oil + v1.Coke_gas + v1.High_pressure_gasification_gas + v1.Pure_coke + v1.Blastfurnace_gas + v1.Converter_gas + v1.Semicoke_gas
	return result
}

func SumEle(v1 *Electricity_And_Heat_Emissions) float64 {
	var result float64
	result = v1.Electricity*co2paifangyingzi[v1.Region] + v1.Heat*0.11
	return result
}

func reflectValSum(val reflect.Value, args ...reflect.Value) reflect.Value {
	kind := val.Kind()
	vi := val.Interface()
	for _, v := range args {
		switch kind {
		case reflect.Int:
			val.Set(reflect.ValueOf(vi.(int) + v.Interface().(int)))
		case reflect.Int8:
			val.Set(reflect.ValueOf(vi.(int8) + v.Interface().(int8)))
		case reflect.Int16:
			val.Set(reflect.ValueOf(vi.(int16) + v.Interface().(int16)))
		case reflect.Int32:
			val.Set(reflect.ValueOf(vi.(int32) + v.Interface().(int32)))
		case reflect.Int64:
			val.Set(reflect.ValueOf(vi.(int64) + v.Interface().(int64)))
		case reflect.Uint:
			val.Set(reflect.ValueOf(vi.(uint) + v.Interface().(uint)))
		case reflect.Uint8:
			val.Set(reflect.ValueOf(vi.(uint8) + v.Interface().(uint8)))
		case reflect.Uint16:
			val.Set(reflect.ValueOf(vi.(uint16) + v.Interface().(uint16)))
		case reflect.Uint32:
			val.Set(reflect.ValueOf(vi.(uint32) + v.Interface().(uint32)))
		case reflect.Uint64:
			val.Set(reflect.ValueOf(vi.(uint64) + v.Interface().(uint64)))
		case reflect.Float32:
			val.Set(reflect.ValueOf(vi.(float32) + v.Interface().(float32)))
		case reflect.Float64:
			val.Set(reflect.ValueOf(vi.(float64) + v.Interface().(float64)))
		}
	}
	return val
}

func StructFieldSum(val interface{}, args ...interface{}) {
	v := reflect.ValueOf(val).Elem()
	t := v.Type()
	num := v.NumField()
	for _, arg := range args {
		vv := reflect.ValueOf(arg).Elem()
		if t != vv.Type() {
			continue
		}
		for i := 0; i < num; i++ {
			//如果是下划线或小写字符开头则忽略
			if t.Field(i).Name[0] >= 95 {
				continue
			}
			reflectValSum(v.Field(i), vv.Field(i))
		}
	}
}

func reflectValMot(val reflect.Value, args ...reflect.Value) reflect.Value {
	kind := val.Kind()
	vi := val.Interface()
	for _, v := range args {
		switch kind {
		case reflect.Int:
			val.Set(reflect.ValueOf(vi.(int) * v.Interface().(int)))
		case reflect.Int8:
			val.Set(reflect.ValueOf(vi.(int8) * v.Interface().(int8)))
		case reflect.Int16:
			val.Set(reflect.ValueOf(vi.(int16) * v.Interface().(int16)))
		case reflect.Int32:
			val.Set(reflect.ValueOf(vi.(int32) * v.Interface().(int32)))
		case reflect.Int64:
			val.Set(reflect.ValueOf(vi.(int64) * v.Interface().(int64)))
		case reflect.Uint:
			val.Set(reflect.ValueOf(vi.(uint) * v.Interface().(uint)))
		case reflect.Uint8:
			val.Set(reflect.ValueOf(vi.(uint8) * v.Interface().(uint8)))
		case reflect.Uint16:
			val.Set(reflect.ValueOf(vi.(uint16) * v.Interface().(uint16)))
		case reflect.Uint32:
			val.Set(reflect.ValueOf(vi.(uint32) * v.Interface().(uint32)))
		case reflect.Uint64:
			val.Set(reflect.ValueOf(vi.(uint64) * v.Interface().(uint64)))
		case reflect.Float32:
			val.Set(reflect.ValueOf(vi.(float32) * v.Interface().(float32)))
		case reflect.Float64:
			val.Set(reflect.ValueOf(vi.(float64) * v.Interface().(float64)))
		}
	}
	return val
}

func StructFieldMot(val interface{}, args ...interface{}) {
	v := reflect.ValueOf(val).Elem()
	t := v.Type()
	num := v.NumField()
	for _, arg := range args {
		vv := reflect.ValueOf(arg).Elem()
		if t != vv.Type() {
			continue
		}
		for i := 0; i < num; i++ {
			//如果是下划线或小写字符开头则忽略
			if t.Field(i).Name[0] >= 95 {
				continue
			}
			reflectValMot(v.Field(i), vv.Field(i))
		}
	}
}
func ReplaceZeroFields(structA, structB interface{}) {
	valA := reflect.ValueOf(structA).Elem()
	valB := reflect.ValueOf(structB).Elem()

	for i := 0; i < valA.NumField(); i++ {
		fieldA := valA.Field(i)
		fieldB := valB.Field(i)

		// 确保两个字段都是 float64 类型
		if fieldA.Kind() == reflect.Float64 && fieldB.Kind() == reflect.Float64 {
			if fieldA.Float() == 0 {
				// 如果 A 结构体中的字段值为 0，则用 B 结构体中对应字段的值替换
				fieldA.SetFloat(fieldB.Float())
			}
		}
	}
}

func EXJuHe(a1, a2, a3, b1, b2, b3, u1, u2, u3 *big.Int) (*big.Int, *big.Int, *big.Int) {
	a := new(big.Int).Add(a1, new(big.Int).Add(a2, a3))
	b := new(big.Int).Add(b1, new(big.Int).Add(b2, b3))
	u := new(big.Int).Mul(u1, new(big.Int).Mul(u2, u3))
	return a, b, u
}

func PKJuHe(r1, r2, r3, s1, s2, s3 *big.Int) (*big.Int, *big.Int) {
	r := new(big.Int).Mul(r1, new(big.Int).Mul(r2, r3))
	s := new(big.Int).Mul(s1, new(big.Int).Mul(s2, s3))
	return r, s
}

func YanZheng(h1, h2, a, b, u, r, s *big.Int) int64 {
	p, _ := new(big.Int).SetString("11354492342642070887128067569374510012280963720199556171757661158988200922003422988172509277025982170385824569298950067555517998425368826215811089098953659", 10)
	g, _ := new(big.Int).SetString("5845106984648640168199956353344722223117328802648400371662308688681439423990963724660129527848629790950521977946161148517314058321656359945838375854931969", 10)
	m1 := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(u, h1, p), r), p)
	m2 := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(u, h2, p), s), p)
	m3 := new(big.Int).Exp(g, a, p)
	m4 := new(big.Int).Exp(g, b, p)
	if m1 == m3 && m2 == m4 {
		return 1
	} else {
		return 0
	}
}
