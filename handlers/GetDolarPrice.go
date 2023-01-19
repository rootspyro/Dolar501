package handlers

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/rootspyro/Dolar501/gen/models"
	"github.com/rootspyro/Dolar501/gen/restapi/operations/dolar"
	"github.com/rootspyro/Dolar501/services"
)

type GetDolarPriceImpl struct {
	srv *services.DolarServices
}

func NewGetDolarPriceImpl(s *services.DolarServices) dolar.GetDolarPriceHandler {
	return &GetDolarPriceImpl{
		srv: s,
	}
}

func(impl *GetDolarPriceImpl)Handle(params dolar.GetDolarPriceParams) middleware.Responder {
	
	strPrice := impl.srv.GetDolarPrice(params.Plataforma)

	if len(strPrice) == 0 {
		errResponse := &models.Default{
			Status: "error",
			Data: fmt.Sprintf("%s usd price not found!", params.Plataforma),
		}

		return dolar.NewGetDolarPriceDefault(404).WithPayload(errResponse)
	}

	price := impl.srv.ParseFloat(strPrice)
	
	data := &models.DolarPrice{
		Plataforma: params.Plataforma,	
		PrecioVes: price,
	}

	response := &models.DolarPriceResponse{
		Status: "success",
		Data: data,
	}

	return dolar.NewGetDolarPriceOK().WithPayload(response)
	
}

