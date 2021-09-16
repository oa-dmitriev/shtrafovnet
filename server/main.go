package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	gw "github.com/oa-dmitriev/shtrafovnet/proto/gen/go"

	"flag"

	"github.com/PuerkitoBio/goquery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	_ "github.com/oa-dmitriev/shtrafovnet/docs"
	"google.golang.org/grpc"
)

var (
	url                = "https://www.rusprofile.ru/search?query=%s&type=ul"
	grpcServerEndpoint = flag.String(
		"grpc-server-endpoint", ":9090", "grpc server endpoint",
	)
)

type LegalInfoFetcher struct {
	gw.UnimplementedLegalInfoFetcherServer
}

func NewLegalInfoFetcher() *LegalInfoFetcher {
	return &LegalInfoFetcher{}
}

func (l *LegalInfoFetcher) GetInfoByInn(
	ctx context.Context,
	inn *gw.Inn,
) (*gw.Info, error) {
	log.Println("Came with inn: ", inn.INN)
	return ParseURL(inn.INN)
}

func ParseURL(inn string) (*gw.Info, error) {
	url := fmt.Sprintf(url, inn)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	info := gw.Info{}
	info.CeoName = doc.Find(".gtm_main_fl").Text()
	info.INN = doc.Find("#clip_inn").Text()
	info.KPP = doc.Find("#clip_kpp").Text()
	doc.Find(".company-name").EachWithBreak(func(i int, s *goquery.Selection) bool {
		info.CompanyName = s.Text()
		return false
	})
	return &info, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	gw.RegisterLegalInfoFetcherServer(server, NewLegalInfoFetcher())

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		server.Serve(lis)
		wg.Done()
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = gw.RegisterLegalInfoFetcherHandlerFromEndpoint(
		ctx, mux, *grpcServerEndpoint, opts,
	)
	if err != nil {
		log.Fatal(err)
	}

	wg.Add(1)
	go func() {
		log.Println("Listening on port: 8081")
		http.ListenAndServe(":8081", mux)
		wg.Done()
	}()

	wg.Wait()
}
