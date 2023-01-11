package response

type ImageClass struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Order int64  `json:"order"`
}

type ImageClassRes struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Order       int64  `json:"order"`
	ImagesCount int64  `json:"images_count"`
}
