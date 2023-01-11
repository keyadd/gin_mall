package request

type ImageClassList struct {
	PageInfo
}

type ImageClassId struct {
	GetById
	PageInfo
}

type CreateImageClass struct {
	Name  string `json:"name"`
	Order int64  `json:"order"`
}
