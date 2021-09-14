package main

import (
	"context"
	"log"
	"net/http"

	"legalinfoclient/legalinfo"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewHandler(client legalinfo.LegalInfoFetcherClient) *handler {
	return &handler{client}
}

type handler struct {
	client legalinfo.LegalInfoFetcherClient
}

type Represent struct {
	INN         string `json:"inn"`
	KPP         string `json:"kpp"`
	CompanyName string `json:"company-name"`
	Name        string `json:"name"`
}

func (h *handler) Index(c *gin.Context) {
	inn := &legalinfo.Inn{INN: "123"}
	info, err := h.client.GetInfoByInn(context.Background(), inn)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, &Represent{
		INN:         info.INN,
		KPP:         info.KPP,
		CompanyName: info.CompanyName,
		Name:        info.CeoName,
	})
}

func main() {
	grpcConn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cannot connect to grpc")
	}
	defer grpcConn.Close()
	client := legalinfo.NewLegalInfoFetcherClient(grpcConn)

	handler := NewHandler(client)
	r := gin.Default()
	r.GET("/", handler.Index)
	r.Run()
}
