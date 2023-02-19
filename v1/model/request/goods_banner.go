package request

type GoodsBannerSet struct {
	GetById
	Url []string `json:"url"`
}
