package request

type Image struct {
	Url          string `json:"url"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
	ImageClassId int64  `json:"image_class_id"`
}

type ImageClassIdRequest struct {
	ImageClassId int64 `json:"image_class_id" form:"image_class_id" binding:"required"`
}

type ImageRename struct {
	GetById
	Name string `json:"name" form:"name" binding:"required"`
}
