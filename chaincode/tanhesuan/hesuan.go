package tanhesuan

import "reflect"

func taoci(hushi *Fossil_Fuel_Combustion, taocizhuanyou *Ceramics_Indusry_Production_Process, dianli *Electricity_And_Heat_Emissions, diqu string) float64 {
	model1 := &Fossil_Fuel_Combustion{
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
	model2 := &Fossil_Fuel_Combustion{
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
	model3 := &Fossil_Fuel_Combustion{
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

	var result float64
	StructFieldMot(model1, model2, model3, hushi)
	a1 := SumFossil(model1) * 44 / 12
	a2 := SumEle(dianli)
	a3 := taocizhuanyou.Material_weight * taocizhuanyou.Material_utilization_ratio * (44/100*taocizhuanyou.CaCO3_content + 44/84*taocizhuanyou.MgCO3_content)
	result = a1 + a2 + a3
	return result
}

func Mayanlian(huashi *Fossil_Fuel_Combustion, ma *Magnesium_smelting_Industry_Production_Process, dianli *Electricity_And_Heat_Emissions) float64 {
	model1 := Fossil_Fuel_Combustion{
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
	model2 := Fossil_Fuel_Combustion{
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
	model3 := Fossil_Fuel_Combustion{
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
	model4 := Magnesium_smelting_Industry_Production_Process{
		Ferrosilicon_yield:   2.79,
		Dolomite_consumption: 0.98,
	}
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
