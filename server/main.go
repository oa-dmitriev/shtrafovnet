package main

import (
	"context"
	"log"
	"net"
	"os"

	"legalinfoserver/legalinfo"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/grpc"
)

var url = "https://www.rusprofile.ru/search?query=%s&type=ul"

type LegalInfoFetcher struct {
	legalinfo.UnimplementedLegalInfoFetcherServer
}

func NewLegalInfoFetcher() *LegalInfoFetcher {
	return &LegalInfoFetcher{}
}

func (l *LegalInfoFetcher) GetInfoByInn(
	ctx context.Context,
	inn *legalinfo.Inn,
) (*legalinfo.Info, error) {
	ParseURL(url)
	return ParseURL(url)
}

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	legalinfo.RegisterLegalInfoFetcherServer(server, NewLegalInfoFetcher())
	server.Serve(lis)
}

func ParseURL(url string) (*legalinfo.Info, error) {
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

	info := legalinfo.Info{}
	info.CeoName = doc.Find(".gtm_main_fl").Text()
	info.INN = doc.Find("#clip_inn").Text()
	info.KPP = doc.Find("#clip_kpp").Text()
	doc.Find(".company-name").EachWithBreak(func(i int, s *goquery.Selection) bool {
		info.CompanyName = s.Text()
		return false
	})
	return &info, nil
}
