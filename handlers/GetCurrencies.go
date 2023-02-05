package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/rootspyro/Dolar501/gen/models"
	"github.com/rootspyro/Dolar501/gen/restapi/operations/dolar"
	"github.com/rootspyro/Dolar501/services"
)

type GetCurrenciesImpl struct {
	srv *services.DolarServices
}

func NewGetCurrenciesImpl( s *services.DolarServices ) dolar.GetDolarCurrenciesHandler {
	return &GetCurrenciesImpl{
		srv: s,
	}
}

func(impl *GetCurrenciesImpl)Handle(params dolar.GetDolarCurrenciesParams, prin interface{}) middleware.Responder {

	currencies, err := impl.srv.GetCurrencyList()

	if err != nil {
		log.Println(err.Error())

		errResponse := &models.Default{
			Status: "error",
			Data: "Algo salio mal!",
		}

		return dolar.NewGetDolarCurrenciesDefault(500).WithPayload(errResponse)
	}

	response := &models.DolarCurrencies{
		Status: "success",
		Data: currencies,
	}

	return dolar.NewGetDolarCurrenciesOK().WithPayload(response) 
}


