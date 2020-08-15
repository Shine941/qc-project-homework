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
	City    string
	CityID  string
	Temp    string
	WD      string
	WS      string
	SD      string
	WSE     string
	Time    string
	IsRadar string
	Radar   string
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
	log.Printf("Results:%v\n", jsonWeather)

	log.Println(jsonWeather.Weatherinfo.City)
	log.Println(jsonWeather.Weatherinfo.WD) //ioutil.WriteFile("wsk101010100.html",input,0644)
}
