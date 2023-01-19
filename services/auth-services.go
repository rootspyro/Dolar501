package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ghAuth "github.com/rootspyro/Dolar501/github"
)

type AuthServices struct {
	ghAuth *ghAuth.GHAuth
}

func NewAuthServices( gh *ghAuth.GHAuth ) *AuthServices{
	return &AuthServices{
		ghAuth: gh,
	}
}

func( s *AuthServices )GetAccessToken(code string) string {
	clientId := s.ghAuth.ClientID
	clientSecret := s.ghAuth.ClientSecret

	body := map[string]string{
		"client_id": clientId,
		"client_secret": clientSecret,
		"code": code,
	}

	requestJSON, err := json.Marshal(body)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	req, reqerr := http.NewRequest(
		"POST",
		s.ghAuth.TokenURL,
		bytes.NewBuffer(requestJSON),
	)

	if reqerr != nil {
		log.Println(err.Error())
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get the response
	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
			log.Panic("Request failed")
	}

	// Response body converted to stringified JSON
	respbody, _ := ioutil.ReadAll(resp.Body)

	// Represents the response received from Github
	type githubAccessTokenResponse struct {
			AccessToken string `json:"access_token"`
			TokenType   string `json:"token_type"`
			Scope       string `json:"scope"`
	}

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	// Return the access token (as the rest of the
	// details are relatively unnecessary for us)
	return ghresp.AccessToken
}
