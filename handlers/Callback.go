package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/rootspyro/Dolar501/gen/models"
	"github.com/rootspyro/Dolar501/gen/restapi/operations/auth"
	ghAuth "github.com/rootspyro/Dolar501/github"
	"github.com/rootspyro/Dolar501/services"
)

type CallbackImpl struct {
	ghAuth *ghAuth.GHAuth
	srv *services.AuthServices
}

func NewCallbackImpl( gh *ghAuth.GHAuth, s *services.AuthServices ) auth.GetAuthTokenHandler {
	return &CallbackImpl{
		ghAuth: gh,
		srv: s,
	}
}

func(impl *CallbackImpl)Handle( params auth.GetAuthTokenParams ) middleware.Responder {

	ghAccessToken:= impl.srv.GetAccessToken(params.Code)	

	if len(ghAccessToken) == 0 {
		errResponse := &models.Default{
			Status: "error",
			Data: "no se pudo obtener el token de acceso!",
		}
		
		return auth.NewGetAuthTokenDefault(500).WithPayload(errResponse)
	}

	okResponse := &models.AccessResponse{
		Status: "success",
		Data: &models.AccessResponseData{
			AccessToken: ghAccessToken,
		},
	}

	return auth.NewGetAuthTokenOK().WithPayload(okResponse)
}

