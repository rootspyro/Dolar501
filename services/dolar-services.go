package services

import (
	"context"
	"fmt"
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

func( impl *DolarServices )GetPlatformsList() []string {
	keys := impl.rdClient.Keys(context.Background(), "*")	
	return keys.Val()
}

func( impl *DolarServices )GetDolarPrice(platform string) string {
	data := impl.rdClient.Get(context.Background(), platform)
	return data.Val()
}

func(impl *DolarServices)ParseFloat(strNum string) float64 {
	
	strNum = strings.Replace(strNum, ",", ".", 2)
	num, _ := strconv.ParseFloat(strNum, 64)
	
	return num
}

func(impl *DolarServices)CalcAverage() (float64, []*models.DolarPrice){

	platforms := []string{
		"monitordolar",
		"enparalelovzla",
	}

	var references []*models.DolarPrice

	var accum float64

	for _, p := range platforms {
		strPrice := impl.GetDolarPrice(p)
		price := impl.ParseFloat(strPrice)
		accum += price

		reference := &models.DolarPrice{
			Plataforma: p,
			PrecioVes: price,
		}
		
		references = append(references, reference)
	}

	average := accum / float64(len(platforms))
	strRounded := fmt.Sprintf("%.2f", average)
	rounded := impl.ParseFloat(strRounded)

	return rounded, references
}
