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

func NewGetPlatformsImpl(s *services.DolarServices ) dolar.GetDolarPlatformsHandler{
	return &GetPlatformsImpl{
		srv: s,
	}
}

func( impl *GetPlatformsImpl )Handle( params dolar.GetDolarPlatformsParams ) middleware.Responder {
	
	data := impl.srv.GetPlatformsList()

	var platformsData []*models.DolarPlatform

	for _, name := range data {
		platform := &models.DolarPlatform{
			Plataforma: name,
			Endpoint: "/api/v1/dolar/" + name,
		}

		platformsData = append(platformsData, platform)
	}

	response := &models.DolarPlatforms{
		Status: "success",
		Data: platformsData,
	}
	return dolar.NewGetDolarPlatformsOK().WithPayload(response)
}


