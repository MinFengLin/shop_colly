package shop

import (
	"github.com/gocolly/colly"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	// "time"
)

type momo_j struct {
	Momo []struct {
			Item string
			Url  string
	} `json:"momo"`
}

func Momo_colly(url string) string {
	momo_c := colly.NewCollector()
	momo_parser_string := ""

	momo_c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})
	// id user #
	momo_c.OnHTML("#osmGoodsName", func(title *colly.HTMLElement) {
		momo_parser_string = momo_parser_string + title.Text
	})
	// class use .
	momo_c.OnHTML(".priceTxtArea", func(price *colly.HTMLElement) {
		momo_parser_string = momo_parser_string + price.Text + "\n"
	})

	momo_c.Visit(url)

	return momo_parser_string
}

// return 要多一個()
// https://opensourcedoc.com/golang-programming/function/ (多個回傳值)
func Momo_parser() string {
	var momo_string = "================Momo================\n"
	filename := "./url.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
	}

	data := momo_j{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
	}

	// Check_iotservice_realtime_status(&data)
	for ii := range data.Momo {
		momo_string = momo_string + "Item: " + data.Momo[ii].Item + "\n"
		momo_string = momo_string + Momo_colly(data.Momo[ii].Url)
	}

	return momo_string
}
