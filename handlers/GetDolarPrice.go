package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-redis/redis/v9"
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

func(impl *GetDolarPriceImpl)Handle(params dolar.GetDolarPriceParams, prin interface{}) middleware.Responder {
	
	result, err := impl.srv.GetDolarPrice(params.Moneda, params.Plataforma)	
	
	if err != nil {
		if err == redis.Nil {
			notFound := &models.Default{
				Status: "error",
				Data: params.Plataforma + " no encontrado!",
			}

			return dolar.NewGetDolarPriceDefault(404).WithPayload(notFound)
		} else {
		
			errResponse := &models.Default{
				Status: "error",
				Data: "Algo malio sal!",
			}
			
			log.Println(err.Error())

			return dolar.NewGetDolarPriceDefault(500).WithPayload(errResponse)
		}
	}

	response := &models.DolarPriceResponse{
		Status: "success",
		Data: &models.DolarPrice{
			Plataforma: params.Plataforma,
			Precio: result,
			Moneda: params.Moneda,
		},
	}

	return dolar.NewGetDolarPriceOK().WithPayload(response)
	
}

