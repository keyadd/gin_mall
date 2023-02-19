package service

import (
	"gin_mall/global"
	"gin_mall/utils"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

func AddImage(files []*multipart.FileHeader, id request.ImageClassIdRequest) []response.Image {
	var imageArr []response.Image

	for _, file := range files {

		fileObj, err := file.Open()
		if err != nil {
			global.GVA_LOG.Error("file.Open()", zap.Error(err))
		}
		uuid := uuid.New()
		key := uuid.String()
		fileExt := strings.ToLower(path.Ext(file.Filename))
		filename := key + fileExt
		ok := utils.UploadFile(global.GVA_CONFIG.Minio.BucketName, filename, fileObj, file.Size)
		if !ok {
			break
		}
		headerUrl := utils.GetFileUrl(global.GVA_CONFIG.Minio.BucketName, filename, time.Second*60)

		if headerUrl == nil {
			continue
		}
		//fmt.Println(headerUrl.String())
		imageUrl := headerUrl.Scheme + "://" + headerUrl.Host + headerUrl.Path
		//fmt.Println(imageUrl)
		image := response.Image{
			Url:          imageUrl,
			Name:         global.GVA_CONFIG.Minio.BucketName + "/" + filename,
			Path:         filename,
			ImageClassId: id.ImageClassId,
			CreateTime:   time.Now().Unix(),
		}

		res, err := model.UploadImage(image)
		//fmt.Println(res)
		imageArr = append(imageArr, res)
	}
	return imageArr
}

func RenameImage(r request.ImageRename) bool {
	res, err := model.RenameImage(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditManager(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteImage(r request.IdsReq) bool {

	var arr []int
	for _, v := range r.Ids {
		arr = append(arr, v)
	}
	res, err := model.DeleteImage(arr)
	if err != nil {
		global.GVA_LOG.Error("model.DeleteImage(arr)", zap.Any("arr", arr), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
