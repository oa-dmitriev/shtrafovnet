package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	gw "github.com/oa-dmitriev/shtrafovnet/proto/gen/go"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/PuerkitoBio/goquery"
	"google.golang.org/grpc"
)

var (
	url                = "https://www.rusprofile.ru/search?query=%s&type=ul"
	grpcServerEndpoint = flag.String(
		"grpc-server-endpoint", "127.0.0.1:9090", "grpc server endpoint",
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
	return ParseURL(url)
}

func ParseURL(url string) (*gw.Info, error) {
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
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// lis, err := net.Listen("tcp", ":8081")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server := grpc.NewServer()

	// legalinfo.RegisterLegalInfoFetcherServer(server, NewLegalInfoFetcher())
	// server.Serve(lis)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterLegalInfoFetcherHandlerFromEndpoint(
		ctx, mux, *grpcServerEndpoint, opts,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	fmt.Println("listening...")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
