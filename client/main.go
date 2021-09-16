package main

import (
	"context"
	"fmt"
	"log"

	gw "github.com/oa-dmitriev/shtrafovnet/proto/gen/go"
	"google.golang.org/grpc"

	"flag"
)

var (
	grpcServerEndpoint = flag.String(
		"grpc-server-endpoint", "127.0.0.1:9090", "grpc server endpoint",
	)
)

func main() {
	flag.Parse()
	grpcConn, err := grpc.Dial(
		*grpcServerEndpoint,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cannot connect to grpc")
	}
	defer grpcConn.Close()
	client := gw.NewLegalInfoFetcherClient(grpcConn)
	info, err := client.GetInfoByInn(context.Background(), &gw.Inn{INN: "123"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ФИО руководителя: %s\nНазвание компании: %s\nИНН: %s\nКПП: %s\n", info.CeoName, info.CompanyName, info.INN, info.KPP)
}
