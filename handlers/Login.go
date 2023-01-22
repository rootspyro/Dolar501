package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rootspyro/Dolar501/gen/restapi/operations/auth"
	ghAuth "github.com/rootspyro/Dolar501/github"
)
type LoginImpl struct {
	ghAuth *ghAuth.GHAuth
}

func NewLoginImpl( gh *ghAuth.GHAuth ) auth.AuthLoginHandler {
	return &LoginImpl{
		ghAuth: gh,
	}
}

func( impl *LoginImpl )Handle(params auth.AuthLoginParams) middleware.Responder {

	r := params.HTTPRequest

	return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {
		url := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s",
		impl.ghAuth.ClientID,
	)

		http.Redirect(w, r, url, 301)
	})

}



