package main

import (
	"context"
	"log"
	"net/http"

	gw "github.com/oa-dmitriev/shtrafovnet/proto/gen/go"

	"github.com/gin-gonic/gin"
)

func NewHandler(client gw.LegalInfoFetcherClient) *handler {
	return &handler{client}
}

type handler struct {
	client gw.LegalInfoFetcherClient
}

type Represent struct {
	INN         string `json:"inn"`
	KPP         string `json:"kpp"`
	CompanyName string `json:"company-name"`
	Name        string `json:"name"`
}

func (h *handler) Index(c *gin.Context) {
	inn := &gw.Inn{INN: "123"}
	info, err := h.client.GetInfoByInn(context.Background(), inn)
	if err != nil {
		log.Fatal("HERE", err)
	}

	c.IndentedJSON(http.StatusOK, &Represent{
		INN:         info.INN,
		KPP:         info.KPP,
		CompanyName: info.CompanyName,
		Name:        info.CeoName,
	})
}

func main() {
	// 	grpcConn, err := grpc.Dial(
	// 		"127.0.0.1:9090",
	// 		grpc.WithInsecure(),
	// 	)
	// 	if err != nil {
	// 		log.Fatalf("cannot connect to grpc")
	// 	}
	// 	defer grpcConn.Close()
	// 	client := legalinfo.NewLegalInfoFetcherClient(grpcConn)
	//
	// 	handler := NewHandler(client)
	// 	r := gin.Default()
	// 	r.GET("/", handler.Index)
	// 	r.Run()
}
