package main

import (
	"flag"
	"fmt"
	"encoding/json"
)

func main() {
	var city,day string

	flag.StringVar(&city, "c", "上海", "城市中文名")
	flag.StringVar(&day, "d", "今天", "可选:今天,昨天,预测")

	flag.Parse()

	body, err := Request(apiUrl + city)
	if err != nil {
		fmt.Println("err was %v", err)
		return
	}

	var r Response
	err = json.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Printf("\nError message: %v", err)
		return
	}
	if r.Status != 200 {
		fmt.Printf("获取天气API出现错误, %s", r.Message)
		return
	}
	Print(day, r)
}
