package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

const apiUrl = "http://www.sojson.com/open/api/weather/json.shtml?city="

func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	return string(body), nil
}

func Print(day string, r Response) {
	fmt.Println("城市:", r.CityName)
	switch day {
	case "今天":
		fmt.Println("湿度:", r.Data.ShiDu)
		fmt.Println("空气质量:", r.Data.Quality)
		fmt.Println("温馨提示:", r.Data.Ganmao)
	case "昨天":
		fmt.Println("日期", r.Data.Yesterday.Date)
		fmt.Println("温度", r.Data.Yesterday.Low, r.Data.Yesterday.High)
		fmt.Println("风量", r.Data.Yesterday.Fx, r.Data.Yesterday.Fl)
		fmt.Println("天气", r.Data.Yesterday.Type)
		fmt.Println("温馨提示", r.Data.Yesterday.Notice)
	case "预测":
		fmt.Println("===========================")
		for _, item := range r.Data.Forecast {
			fmt.Println("日期", item.Date)
			fmt.Println("温度", item.Low, item.High)
			fmt.Println("风量", item.Fx, item.Fl)
			fmt.Println("天气", item.Type)
			fmt.Println("温馨提示", item.Notice)
			fmt.Println("===========================")
		}
	default:
		fmt.Println("大熊你是想刁难我胖虎吗?_?")

	}
}