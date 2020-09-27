package java

// 查询货币配置列表
type GetCurrencyListStruct struct {
	Message string                      `json:"message"`
	RspCode string                      `json:"rspCode"`
	RspTime int                         `json:"rspTime"`
	Data    []GetCurrencyListDataStruct `json:"data"`
}

//"currency": "USDT",
//"sort": 1,
//"degree": 2,
//"hide": 0,
//"bettingMax": 1000.000000,
//"bettingMin": 1.000000,
//"operatingRange": 1.000000,
//"imageUrl": "/public-test/img/coin/15905807050632UL8JD1J.png",
//"svgUrl": "/public-test/img/coin/1596112099855SHWDDR1008.svg"
type GetCurrencyListDataStruct struct {
	Currency       string  `json:"currency"`
	Sort           int     `json:"sort"`
	Degree         int     `json:"degree"`
	Hide           int     `json:"hide"`
	BettingMax     float64 `json:"bettingMax"`
	BettingMin     float64 `json:"bettingMin"`
	OperatingRange float64 `json:"operatingRange"`
	ImageUrl       string  `json:"imageUrl"`
	SvgUrl         string  `json:"svgUrl"`
}

// 查询货币配置列表
type GetBalanceListStruct struct {
	Message string                              `json:"message"`
	RspCode string                              `json:"rspCode"`
	RspTime int                                 `json:"rspTime"`
	Data    map[string]GetBalanceListDataStruct `json:"data"`
}

// 查询货币配置列表
type GetBalanceListDataStruct struct {
	Total          float64 `json:"total"`
	TotalAvailable float64 `json:"totalAvailable"`
	Available      float64 `json:"available"`
	Frozen         float64 `json:"frozen"`
	WorthUSDT      float64 `json:"worthUSDT"`
	Sort           int     `json:"sort"`
}
