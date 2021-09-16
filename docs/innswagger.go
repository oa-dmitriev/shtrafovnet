package docs

import (
	gw "github.com/oa-dmitriev/shtrafovnet/proto/gen/go"
)

// swagger:route POST /v1/inn{INN} inn-tag idOfInnEndpoint
// Get legal info by INN
// responses:
//   200: infoResponse

// swagger:response infoResponse
type innResponseWrapper struct {
	// in:body
	Body gw.Info
}

// swagger:parameters idOfInnEndpoint
type innParamsWrapper struct {
	// in:body
	Body gw.Inn
}
