package response

type UserLevel struct {
	Id       int64  `json:"id" `
	Name     string `json:"name"`
	Level    int64  `json:"level"`
	Status   int64  `json:"status"`
	Discount int64  `json:"discount"`
	MaxPrice int64  `json:"max_price"`
	MaxTimes int64  `json:"max_times"`
}

type UserLevelList struct {
	Id   int64  `json:"id" `
	Name string `json:"name"`
}

type CreateUserLevel struct {
	Id       int64  `json:"id" `
	Name     string `json:"name"`
	Level    int64  `json:"level"`
	Status   int64  `json:"status"`
	Discount int64  `json:"discount"`
	MaxPrice int64  `json:"max_price"`
	MaxTimes int64  `json:"max_times"`
}
