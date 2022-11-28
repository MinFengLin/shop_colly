package crawler

import (
	"github.com/gocolly/colly"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Momo_colly(url string, item string, target_price int) string {
	momo_c := colly.NewCollector()
	momo_parser_string := "Item: " + item + "\n"
	var parser_price int

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
		re := regexp.MustCompile("[0-9]+")
		parser_price, _ = strconv.Atoi(strings.Join(re.FindAllString(price.Text, -1), ""))
	})

	_ = momo_c.Visit(url)
	momo_parser_string = momo_parser_string + "Go to link: 🔗 " + url + "\n"

	if parser_price <= target_price {
		momo_parser_string = "✔ 快去搶購:\n目標價 -> " + strconv.Itoa(target_price) + "\n現價 -> " + strconv.Itoa(parser_price) + "\n" + momo_parser_string
	} else {
		momo_parser_string = ""
	}

	return momo_parser_string
}

func Momo_parser() momo_j {
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

	return data
}

func Momo_list_data() string {
	data := Momo_parser()
	var list = "-\n"

	for ii := range data.Momo {
		list = list + data.Momo[ii].Item + "\n -> 目標價格：" + data.Momo[ii].Target_price + "\n 網址-(" + data.Momo[ii].Url + ")" + "\n"
	}
	list = list + "-\n"

	return list
}

func Momo_parser_data() string {
	var momo_string = ""
	data := Momo_parser()

	for ii := range data.Momo {
		target_price, _ := strconv.Atoi(data.Momo[ii].Target_price)
		momo_string = momo_string + Momo_colly(data.Momo[ii].Url, data.Momo[ii].Item, target_price)
	}

	if len(momo_string) > 0 {
		return momo_string
	} else {
		return "😔 追蹤的商品,皆為高於目標價格"
	}
}
