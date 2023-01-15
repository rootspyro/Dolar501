package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/rootspyro/Dollar-VzlAPI/gen/models"
	"github.com/rootspyro/Dollar-VzlAPI/gen/restapi/operations/dolar"
	"github.com/rootspyro/Dollar-VzlAPI/services"
)

type GetDolarAverageImpl struct {
	srv *services.DolarServices
}

func NewGetDolarAverageImpl(s *services.DolarServices) dolar.GetDolarAverageHandler {
	return &GetDolarAverageImpl{
		srv: s,
	}
}

func(impl *GetDolarAverageImpl)Handle(params dolar.GetDolarAverageParams) middleware.Responder{

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
