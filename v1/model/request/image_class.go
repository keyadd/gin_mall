package request

type ImageClassList struct {
	PageInfo
}

type ImageClassId struct {
	GetById
	PageInfo
}

type CreateImageClass struct {
	Name  string `json:"name" form:"name"`
	Order int64  `json:"order" form:"order"`
}

type EditImageClass struct {
	GetById
	Name  string `json:"name" form:"name"`
	Order int64  `json:"order" form:"order"`
}
