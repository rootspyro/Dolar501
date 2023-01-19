package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/rootspyro/Dolar501/gen/models"
	"github.com/rootspyro/Dolar501/gen/restapi/operations/dolar"
	"github.com/rootspyro/Dolar501/services"
)

type GetDolarAverageImpl struct {
	srv *services.DolarServices
}

func NewGetDolarAverageImpl(s *services.DolarServices) dolar.GetDolarAverageHandler {
	return &GetDolarAverageImpl{
		srv: s,
	}
}

func(impl *GetDolarAverageImpl)Handle(params dolar.GetDolarAverageParams, prin interface{}) middleware.Responder{

	dollar, references := impl.srv.CalcAverage()


	data := &models.DolarAverage{
		DolarParalelo: dollar,
		Referencias: references,
	}

	response := &models.DolarAverageResponse{
		Status: "success",
		Data: data,
	}
	
	return dolar.NewGetDolarAverageOK().WithPayload(response)
}
