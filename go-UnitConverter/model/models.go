package model

type Request struct {
	Input    float64 `json:"input"`
	UnitFrom string  `json:"unitForm"`
	UnitTo   string  `json:"unitTo"`
}

type Response struct {
	Result float64 `json:"result"`
	Unit   string  `json:"string"`
}
