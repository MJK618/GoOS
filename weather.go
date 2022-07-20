package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io"
	"io/ioutil"
	"net/http"
)

func showWeatherApp(w fyne.Window) {
	//a := app.New()
	//w := a.NewWindow("GoWeather")

	//Call API
	city := "Delhi"
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=3e7570272b8a79c70b7cb0b971d78319")

	if err != nil {
		fmt.Println(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	weather, err := UnmarshalGoWeather(body)
	if err!= nil{
		fmt.Println(err)
	}
	imageGoWeatherApp := canvas.NewImageFromFile("./systemfiles/weather/img/weather.jpg")
	imageGoWeatherApp.FillMode = canvas.ImageFillOriginal
	label1 := widget.NewLabel("Weather for " + city)
	label2 := canvas.NewText(fmt.Sprintf("Country: %s", weather.Sys.Country), color.White)
	label3 := canvas.NewText(fmt.Sprintf("Temprature: %.2f", weather.Main.Temp), color.White)
	label4 := canvas.NewText(fmt.Sprintf("Humidity: %d", weather.Main.Humidity), color.White)
	label5 := canvas.NewText(fmt.Sprintf("Wind Speed: %.2f", weather.Wind.Speed), color.White)
	label6 := canvas.NewText(fmt.Sprintf("Clouds %d", weather.Clouds.All), color.White)
	label7 := canvas.NewText(fmt.Sprintf("Sunrise: %d", weather.Sys.Sunrise), color.White)
	label8 := canvas.NewText(fmt.Sprintf("Sunset: %d", weather.Sys.Sunset), color.White)
	//combo := widget.NewSelect([]string{"Delhi", "Mumbai", "Noida", "Gurgaon"}, func(value string) {
	//	label1.SetText(value)
	//	city = value
	//})
	weatherContainer := container.NewVBox(
			imageGoWeatherApp,
			//combo,
			label1,
			label2,
			label3,
			label4,
			label5,
			label6,
			label7,
			label8,
		)
    w.SetContent(container.NewBorder(nil,nil,appPanel,nil,weatherContainer))
	//w.Resize(fyne.NewSize(250,250))
	//w.ShowAndRun()
	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    goWeather, err := UnmarshalGoWeather(bytes)
//    bytes, err = goWeather.Marshal()


func UnmarshalGoWeather(data []byte) (GoWeather, error) {
	var r GoWeather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GoWeather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GoWeather struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
