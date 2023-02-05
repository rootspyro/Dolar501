package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/rootspyro/Dolar501/gen/models"
	"github.com/rootspyro/Dolar501/gen/restapi/operations/dolar"
	"github.com/rootspyro/Dolar501/services"
)

type GetPlatformsImpl struct {
	srv *services.DolarServices
}

func NewGetPlatformsImpl( s *services.DolarServices ) dolar.GetCurrencyPlatformsHandler {
	return &GetPlatformsImpl{
		srv: s,
	}
}

func(impl *GetPlatformsImpl)Handle(params dolar.GetCurrencyPlatformsParams, princ interface{}) middleware.Responder {

	result, err := impl.srv.GetCurrencyPlatforms(params.Moneda)

	if err != nil {
		errResponse := &models.Default{
			Status: "error",
			Data: "Algo salio mal!",
		}

		return dolar.NewGetCurrencyPlatformsDefault(500).WithPayload(errResponse)
	}

	if result == nil {
		notFound := &models.Default{
			Status: "error",
			Data: params.Moneda + " no fue encontrada!",
		}

		return dolar.NewGetCurrencyPlatformsDefault(404).WithPayload(notFound)
	}

	response := &models.DolarPlatforms{
		Status: "success",
		Data: result,
	} 

	return dolar.NewGetCurrencyPlatformsOK().WithPayload(response)

}
