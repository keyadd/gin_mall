package response

type Image struct {
	Id           int64  `json:"id"`
	Url          string `json:"url"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
	ImageClassId int64  `json:"image_class_id"`
}
