package services

import (
	"context"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v9"
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
