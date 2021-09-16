package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
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
	log.Println("Came with inn: ", inn.INN)
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

	// -----

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

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	wg.Add(1)
	go func() {
		http.ListenAndServe(":8081", mux)
		wg.Done()
	}()

	wg.Wait()
}
