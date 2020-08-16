package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WeatherInfoJson struct {
	Weatherinfo WeatherinfoObject
}

type WeatherinfoObject struct {
	City string
	Temp string
	WD   string
	WS   string
	SD   string
	AP   string
	Time string
	Sm   string
}

func main() {
	var num string
	var buf bytes.Buffer
	fmt.Println("请输入所在城市的Id号")
	fmt.Scanln(&num)
	buf.WriteString("http://www.weather.com.cn/data/sk/")
	buf.WriteString(num)
	buf.WriteString(".html")
	num = buf.String()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	resp, err := http.Get(num)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	input, err := ioutil.ReadAll(resp.Body)

	var jsonWeather WeatherInfoJson
	json.Unmarshal(input, &jsonWeather)
	log.Printf("详细结果:\n%v\n", jsonWeather)
	fmt.Println("大致情况如下")
	fmt.Println("城市:", jsonWeather.Weatherinfo.City)
	fmt.Println("时间:", jsonWeather.Weatherinfo.Time)
	fmt.Println("温度:", jsonWeather.Weatherinfo.Temp)
	fmt.Println("风向:", jsonWeather.Weatherinfo.WD)
	fmt.Println("风速:", jsonWeather.Weatherinfo.WS)
	fmt.Println("空气湿度:", jsonWeather.Weatherinfo.SD)
	fmt.Println("sm指数:", jsonWeather.Weatherinfo.Sm)
	fmt.Println("大气压:", jsonWeather.Weatherinfo.WS)
}
