package myconstants


const(  //模仿枚举
	MIN_1 = iota
	MIN_5
	MIN_15
	MIN_30
	MIN_60
	DAY_1
	MON_1
	WEEK_1
	YEAR_1
)
func GetKlinePeriod(enumID int) string {
	Period := []string{"1min", "5min", "15min",
			"30min","60min", "1day",
			"1mon", "1week", "1year" }
	return Period[enumID]
}

const (
	percent_10 = iota
	step_1
	step_2
	step_3
	step_4
	step_5
)
func GetDepthType(enumID int)string{
	depthType := []string{ "percent10", "step1", "step2", "step3", "step4", "step5" }
	return depthType[enumID]
}

const (
	BTC_SYMBOL = iota
	LTC_SYMBOL
	ETH_SYMBOL
)

func GetSymbol(enumID int)string{
	symbols := []string{ "btccny", "ltccny","ethcny"  }
	return symbols[enumID]
}
