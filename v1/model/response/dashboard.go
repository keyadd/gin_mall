package response

type Statistics1 struct {
	Title     string `json:"title"`
	Value     string `json:"value"`
	Unit      string `json:"unit"`
	UnitColor string `json:"unit_color"`
	SubTitle  string `json:"sub_title"`
	SubValue  string `json:"sub_value"`
	SubUnit   string `json:"sub_unit"`
}

type Statistics2 struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

type Result struct {
	Goods []Statistics2 `json:"goods"`
	Order []Statistics2 `json:"order"`
}

type Res struct {
	Time string `json:"time"`
	Num  string `json:"num"`
}

type ListRes struct {
	X []string `json:"x"`
	Y []int64  `json:"y"`
}
