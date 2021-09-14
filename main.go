package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Info struct {
	Name    string
	CeoName string
	INN     string
	KPP     string
}

var url = "https://www.rusprofile.ru/search?query=%s&type=ul"

func main() {
	ParseURL(url)
}

func ParseURL(url string) (*Info, error) {
	file, err := os.Open("text.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// b, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(b))
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Find(".gtm_main_fl").Text())
	fmt.Println(doc.Find("#clip_inn").Text())
	fmt.Println(doc.Find("#clip_kpp").Text())

	doc.Find(".company-name").EachWithBreak(func(i int, s *goquery.Selection) bool {
		fmt.Println(s.Text())
		return false
	})

	return nil, nil
}
