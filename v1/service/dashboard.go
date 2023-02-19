package service

import (
	"gin_mall/v1/model"
	"gin_mall/v1/model/request"
	"gin_mall/v1/model/response"
	"strconv"
	"time"
)

func GetStatistics1() response.PageResult {
	//总订单量
	totalOrder := model.GetOrderCount()

	totalSale := model.GetTotalSale()

	totalUser := model.GetUserCount()
	totalPay := model.GetOrderPayCount()

	totalPayStr := strconv.Itoa(int(totalPay))

	totalOrderStr := strconv.Itoa(int(totalOrder))
	//fmt.Println(totalPay)

	totalUserStr := strconv.Itoa(int(totalUser))

	totalSaleString := strconv.FormatFloat(totalSale, 'f', 10, 32)
	list := []response.Statistics1{
		{
			Title:     "支付订单",
			Value:     totalPayStr,
			Unit:      "年",
			UnitColor: "success",
			SubTitle:  "总支付订单",
			SubValue:  totalPayStr,
			SubUnit:   "",
		}, {
			Title:     "订单量",
			Value:     totalOrderStr,
			Unit:      "周",
			UnitColor: "danger",
			SubTitle:  "转化率",
			SubValue:  "60%",
			SubUnit:   "",
		}, {
			Title:     "销售额",
			Value:     totalSaleString,
			Unit:      "年",
			UnitColor: "danger",
			SubTitle:  "总销售额",
			SubValue:  totalSaleString,
			SubUnit:   "",
		}, {
			Title:     "新增用户",
			Value:     totalUserStr,
			Unit:      "年",
			UnitColor: "warning",
			SubTitle:  "总用户",
			SubValue:  totalUserStr,
			SubUnit:   "人",
		},
	}
	result := response.PageResult{
		List: list,
	}

	return result

}

func GetStatistics2() response.PageResult {

	//checking :=

	// 审核中
	checking := model.GetChecking()
	// 销售中
	saling := model.GetSaling()
	// 已下架
	off := model.GetOff()
	//库存预警
	min_stock := model.GetMinStock()

	// 待付款
	nopay := model.GetOrderNopayCount()

	// 待发货
	noship := model.GetOrderNoshipCount()
	// 已发货

	shiped := model.GetOrderShipedCount()

	// 退款中

	refunding := model.GetOrderRefundingCount()

	goods := []response.Statistics2{
		{
			Label: "审核中",
			Value: checking,
		}, {
			Label: "销售中",
			Value: saling,
		},
		{
			Label: "已下架",
			Value: off,
		},
		{
			Label: "库存预警",
			Value: min_stock,
		},
	}

	order := []response.Statistics2{
		{
			Label: "待付款",
			Value: nopay,
		}, {
			Label: "待发货",
			Value: noship,
		},
		{
			Label: "已发货",
			Value: shiped,
		},
		{
			Label: "退款中",
			Value: refunding,
		},
	}
	r := response.Result{
		Goods: goods,
		Order: order,
	}
	result := response.PageResult{
		List: r,
	}
	return result
}

func GetStatistics3(r request.Statistics3) (u response.PageResult) {
	//now := time.Now()
	d := "2021-07-24"

	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, d, loc)
	endTime := theTime.Add(time.Hour * 24).Add(time.Minute * 59)

	//fmt.Println(theTime)
	//r.DateType = "hour"
	var list response.ListRes

	if r.DateType == "month" {

		duration, _ := time.ParseDuration("-720h")
		startTime := endTime.Add(duration)
		//fmt.Println(startTime)
		dateStr := "%m-%d"
		//fmt.Println(startTime.Format("01-02 15:04:05"))
		//dur, _ := time.ParseDuration("24h")
		var date []string
		for i := 0; i < 30; i++ {

			add := startTime.AddDate(0, 0, i)
			date = append(date, add.Format("01-02"))

		}

		var DateRes []response.Res
		for _, j := range date {
			r1 := response.Res{
				Time: j,
				Num:  "0",
			}
			DateRes = append(DateRes, r1)
		}
		res, _ := model.GetOrder(startTime.Unix(), endTime.Unix(), dateStr)
		for _, i := range DateRes {
			for _, v := range res {
				if v.Time == i.Time {
					i.Num = v.Num
				}

			}

			parseInt, _ := strconv.ParseInt(i.Num, 10, 64)
			list.X = append(list.X, i.Time)
			list.Y = append(list.Y, parseInt)
		}

	} else if r.DateType == "week" {
		duration, _ := time.ParseDuration("-168h")
		startTime := endTime.Add(duration)
		dateStr := "%m-%d"
		var date []string
		for i := 0; i < 7; i++ {
			add := startTime.AddDate(0, 0, i)
			date = append(date, add.Format("01-02"))

		}

		var DateRes []response.Res
		for _, j := range date {
			r1 := response.Res{
				Time: j,
				Num:  "0",
			}
			DateRes = append(DateRes, r1)
		}
		res, _ := model.GetOrder(startTime.Unix(), endTime.Unix(), dateStr)

		for _, i := range DateRes {
			for _, v := range res {
				if v.Time == i.Time {
					i.Num = v.Num
				}

			}

			parseInt, _ := strconv.ParseInt(i.Num, 10, 64)
			list.X = append(list.X, i.Time)
			list.Y = append(list.Y, parseInt)
		}

	} else if r.DateType == "hour" {

		duration, _ := time.ParseDuration("-24h")
		startTime := endTime.Add(duration)
		dateStr := "%H"
		var times []string
		for i := 0; i < 24; i++ {
			s := strconv.Itoa(i) + "h"
			dur, _ := time.ParseDuration(s)
			times = append(times, endTime.Add(dur).Format("15"))
		}

		var DateRes []response.Res
		for _, j := range times {
			r1 := response.Res{
				Time: j,
				Num:  "0",
			}
			DateRes = append(DateRes, r1)
		}

		res, _ := model.GetOrder(startTime.Unix(), endTime.Unix(), dateStr)

		for _, i := range DateRes {
			for _, v := range res {
				if v.Time == i.Time {
					i.Num = v.Num
				}

			}

			parseInt, _ := strconv.ParseInt(i.Num, 10, 64)
			list.X = append(list.X, i.Time)
			list.Y = append(list.Y, parseInt)
		}
	}
	result := response.PageResult{
		List: list,
	}
	return result
}
