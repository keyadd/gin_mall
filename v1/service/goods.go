package service

import (
	"encoding/json"
	"fmt"
	"gin_mall/global"
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// GetGoodsList 获取Goods 用户列表
func GetGoodsList(list request.GoodsList) (res response.GoodsListRes) {
	GoodsList, total, err := model.GetGoodsList(list)
	if err != nil {
		global.GVA_LOG.Error("model.GetGoodsList(list)", zap.Any("list", list), zap.Error(err))
		return
	}
	var goodsResult []*response.GoodsList
	for _, i := range GoodsList {

		UpdateTime, _ := strconv.Atoi(i.UpdateTime)
		unixUpdateTime := time.Unix(int64(UpdateTime), 0)
		i.UpdateTime = unixUpdateTime.Format("2006-01-02 15:04:05")

		createTime, _ := strconv.Atoi(i.CreateTime)
		unixCreateTime := time.Unix(int64(createTime), 0)
		i.CreateTime = unixCreateTime.Format("2006-01-02 15:04:05")

		if i.SkuType == 1 {
			var skus []response.SkusInfo
			if len(i.GoodsSkus) != 0 {
				for _, j := range i.GoodsSkus {
					_ = json.Unmarshal([]byte(j.Skus), &skus)
				}
			}
			i.SkusInfo = skus
		}
		goodsResult = append(goodsResult, i)
	}

	categoryList := GetCategoryList()
	listRes := response.GoodsListRes{
		List:     goodsResult,
		Total:    total,
		Category: categoryList.List,
	}

	return listRes
}

// CreateGoods 创建Goods用户
func CreateGoods(r request.AddGoods) (res response.GoodsCreate) {
	Goods := response.GoodsCreate{
		Title:        r.Title,
		CategoryId:   r.CategoryId,
		CreateTime:   time.Now().Unix(),
		Cover:        r.Cover,
		Desc:         r.Desc,
		Unit:         r.Unit,
		Stock:        r.Stock,
		MinStock:     r.MinStock,
		Status:       r.Status,
		StockDisplay: r.StockDisplay,
		MinOprice:    r.MinOprice,
		MinPrice:     r.MinPrice,
	}

	u, err := model.CreateGoods(Goods)

	if err != nil {
		global.GVA_LOG.Error("model.CreateGoods(Goods)", zap.Any("Goods", Goods), zap.Error(err))
		return
	}
	return u
}

func EditGoods(r request.EditGoods) bool {
	r.UpdateTime = time.Now().Unix()

	res, err := model.EditGoods(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditGoods(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func DeleteAllGoods(r request.IdsReq) bool {
	var arr []int
	for _, v := range r.Ids {
		arr = append(arr, v)
	}
	res, err := model.DeleteAllGoodsId(arr)
	if err != nil {
		global.GVA_LOG.Error("model.EditGoods(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func DestroyAllGoods(r request.IdsReq) bool {
	var arr []int
	for _, v := range r.Ids {
		arr = append(arr, v)
	}
	res, err := model.DestroyGoodsId(arr)
	if err != nil {
		global.GVA_LOG.Error("model.EditGoods(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}

}

func ChangeStatusGoods(r request.StatusIdsRequest) bool {
	//var arr []int
	//for _, v := range r.Ids {
	//	arr = append(arr, v)
	//}
	res, err := model.ChangeStatusGoodsId(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditGoods(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}
func GoodsCheck(r request.GoodsCheck) bool {
	res, err := model.GoodsCheck(r)
	if err != nil {
		global.GVA_LOG.Error("model.EditGoods(r)", zap.Any("r", r), zap.Error(err))
		return false
	}
	if res == 1 {
		return true
	} else {
		return false
	}
}

func GoodsRead(r request.GetById) (result response.GoodsRead) {
	//fmt.Println(r.Id)
	goodsRes, err := model.GetGoodsId(r)
	if err != nil {
		global.GVA_LOG.Error("model.GetGoodsId(r)", zap.Any("r", r), zap.Error(err))
		return
	}
	//fmt.Println(goodsRes)
	banners, err := model.GetBannerId(goodsRes.Id)
	if err != nil {
		global.GVA_LOG.Error("model.GetBannerId(goodsRes.Id)", zap.Any("goodsRes.Id", goodsRes.Id), zap.Error(err))
		return
	}
	goodsRes.GoodsBanner = banners
	if goodsRes.SkuType == 1 {

		skusArr, err := model.GetGoodsSkusReadId(goodsRes.Id)
		if err != nil {
			global.GVA_LOG.Error("model.GetGoodsSkusId(goodsRes.Id)", zap.Any("goodsRes.Id", goodsRes.Id), zap.Error(err))
			return
		}

		res, err := model.GetGoodsSkusJson(goodsRes.Id)

		if err != nil {
			global.GVA_LOG.Error("model.GetGoodsSkusId(goodsRes.Id)", zap.Any("goodsRes.Id", goodsRes.Id), zap.Error(err))
			return
		}
		var skus []response.SkusInfo
		for _, k := range res {
			_ = json.Unmarshal([]byte(k.Skus), &skus)
		}
		fmt.Println(skus)
		skuscard, err := model.GetGoodsSkusCardId(goodsRes.Id)
		if err != nil {
			global.GVA_LOG.Error("model.GetGoodsSkusId(goodsRes.Id)", zap.Any("goodsRes.Id", goodsRes.Id), zap.Error(err))
			return
		}
		//fmt.Println(skuscard)
		var goodsSkusCard []response.GoodsSkusCard
		if len(skuscard) != 0 {
			for _, s := range skuscard {
				res, err := model.GetGoodsSkusCardValueId(s.Id)
				if err != nil {
					global.GVA_LOG.Error("model.GetGoodsSkusId(goodsRes.Id)", zap.Any("goodsRes.Id", goodsRes.Id), zap.Error(err))
					return
				}
				s.GoodsSkusCardValue = res
				goodsSkusCard = append(goodsSkusCard, s)
			}
		}

		goodsRes.GoodsSkus = skusArr
		if len(skuscard) != 0 {
			goodsRes.GoodsSkusCard = goodsSkusCard
		} else {
			goodsRes.GoodsSkusCard = skuscard
		}

	}

	return goodsRes

}

func GoodsBannerSet(r request.GoodsBannerSet) (result []response.GoodsBanner) {

	var arr []response.GoodsBanner
	for _, v := range r.Url {
		banner := response.GoodsBanner{
			GoodsId:    r.Id,
			Url:        v,
			CreateTime: time.Now().Unix(),
		}
		arr = append(arr, banner)
	}
	res, err := model.GoodsBannerSet(arr)
	if err != nil {
		global.GVA_LOG.Error("model.GoodsBannerSet(arr)", zap.Any("arr", arr), zap.Error(err))
		return
	}
	return res

}

func GoodsUpdateSkus(r request.GoodsUpdateSkus) bool {
	//
	if r.SkuType == 0 {
		SkuValue, _ := json.Marshal(r.SkuValue)
		res, err := model.GoodsUpdateSkuType(r)
		if err != nil {
			global.GVA_LOG.Error("model.GoodsUpdateSkuType(r)", zap.Any("r", r), zap.Error(err))
			return false
		}
		res, err = model.GoodsUpdateSkus(r, string(SkuValue))
		if err != nil {
			global.GVA_LOG.Error(" model.GoodsUpdateSkus(r, string(SkuValue))", zap.Any("r", r), zap.Error(err))
			return false
		}
		if res == 1 {
			return true
		} else {
			return false
		}

	} else {

		var skus []response.UpdateGoodsSkus
		for _, v := range r.GoodsSkus {
			marshal, _ := json.Marshal(v.Skus)
			pprice, _ := strconv.ParseFloat(v.Pprice, 64)
			oprice, _ := strconv.ParseFloat(v.Oprice, 64)
			cprice, _ := strconv.ParseFloat(v.Cprice, 64)
			goodsSkus := response.UpdateGoodsSkus{
				Skus:    string(marshal),
				Id:      v.Id,
				GoodsId: v.GoodsId,
				Image:   v.Image,
				Pprice:  pprice,
				Oprice:  oprice,
				Cprice:  cprice,
				Stock:   v.Stock,
				Volume:  v.Volume,
				Weight:  v.Weight,
				Code:    v.Code,
			}
			skus = append(skus, goodsSkus)
		}
		_, err := model.GoodsUpdateSkuType(r)
		if err != nil {
			global.GVA_LOG.Error("model.GoodsUpdateSkuType(r)", zap.Any("r", r), zap.Error(err))
			return false
		}

		err = model.GoodsUpdateSkusArr(skus)
		if err != nil {
			global.GVA_LOG.Error("model.GoodsUpdateSkusArr(skus)", zap.Any("skus", skus), zap.Error(err))
			return false
		}
		return true

	}

}
