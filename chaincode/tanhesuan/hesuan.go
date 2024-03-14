package tanhesuan

import "reflect"

func Taoci(hushi *Fossil_Fuel_Combustion, taocizhuanyou *Ceramics_Indusry_Production_Process, dianli *Electricity_And_Heat_Emissions, model1 Fossil_Fuel_Combustion, model2 Fossil_Fuel_Combustion, model3 Fossil_Fuel_Combustion) float64 {
	var result float64
	StructFieldMot(&model1, &model2, &model3, hushi)
	a1 := SumFossil(&model1) * 44 / 12
	a2 := SumEle(dianli)
	a3 := taocizhuanyou.Material_weight * taocizhuanyou.Material_utilization_ratio * (44/100*taocizhuanyou.CaCO3_content + 44/84*taocizhuanyou.MgCO3_content)
	result = a1 + a2 + a3
	return result
}

func Mayanlian(huashi *Fossil_Fuel_Combustion, ma *Magnesium_smelting_Industry_Production_Process, dianli *Electricity_And_Heat_Emissions, model1 Fossil_Fuel_Combustion, model2 Fossil_Fuel_Combustion, model3 Fossil_Fuel_Combustion, model4 Magnesium_smelting_Industry_Production_Process) float64 {
	var result float64
	StructFieldMot(&model1, &model2, &model3, huashi)
	a1 := SumFossil(&model1) * 44 / 12
	a2 := SumEle(dianli)
	StructFieldMot(&model4, ma)
	a3 := model4.Ferrosilicon_yield + model4.Dolomite_consumption*0.478
	result = a1 + a2 + a3
	return result
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
