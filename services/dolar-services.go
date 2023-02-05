package services

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/rootspyro/Dolar501/gen/models"
)

type DolarServices struct {
	rdClient *redis.Client
}

func NewDolarServices( client *redis.Client ) *DolarServices {
	return &DolarServices{
		rdClient: client,
	}
}

func( s *DolarServices )GetCurrencyList() ([]*models.DolarCurrency, error) {

	var currencies []*models.DolarCurrency

	keys, err := s.rdClient.Keys(context.Background(), "*").Result()

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		currency := models.DolarCurrency{
			Moneda: key,
			Endpoint: "/api/v1/dolar/" + key,
		}

		currencies = append(currencies, &currency)
	}
	
	return currencies, err
}

func( s *DolarServices )GetPlatformsList() []string {
	keys := s.rdClient.Keys(context.Background(), "*")	
	return keys.Val()
}

func( s *DolarServices )GetCurrencyPlatforms( currency string ) ([]*models.DolarPlatform, error) {

	var platforms []*models.DolarPlatform

	result, err := s.rdClient.HGetAll(context.Background(), currency).Result()

	if err != nil {
		return nil, err
	}

	for pName, price := range result {
		platform := models.DolarPlatform {
			Plataforma: pName,
			Precio: s.ParseFloat(price),
			Endpoint: "/api/v1/dolar/" + currency + "/" + pName,
		}

		platforms = append(platforms, &platform)
	}

	return platforms, nil 

}

func( s *DolarServices )GetDolarPrice(currency, platform string) (float64, error) {
	result, err := s.rdClient.HGet(context.Background(), currency, platform).Result()

	if err != nil {
		return 0, err
	}

	price := s.ParseFloat(result)
	return price, nil 

}

func(s *DolarServices)ParseFloat(strNum string) float64 {
	
	strNum = strings.Replace(strNum, ",", ".", 2)
	num, _ := strconv.ParseFloat(strNum, 64)
	
	return num
}

